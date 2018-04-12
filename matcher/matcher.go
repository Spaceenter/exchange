// Package matcher is an engine that receives and matches orders.
package matcher

import (
	"errors"

	"github.com/golang/protobuf/ptypes"
	tspb "github.com/golang/protobuf/ptypes/timestamp"
	"github.com/google/btree"
	pb "github.com/spaceenter/exchange/proto"
)

type Matcher struct {
	tradingPair pb.TradingPair
	askTree     *btree.BTree
	bidTree     *btree.BTree
}

func New(tradingPair pb.TradingPair) *Matcher {
	return &Matcher{
		tradingPair: tradingPair,
		askTree:     btree.New(2),
		bidTree:     btree.New(2),
	}
}

func (m *Matcher) SubmitOrder(order pb.Order) (tradeEvents []*pb.TradeEvent,
	orderBookEvents []*pb.OrderBookEvent, err error) {
	timestamp, err := ptypes.Timestamp(order.Timestamp)
	if err != nil {
		return
	}
	item := orderItem{
		orderId:   order.OrderId,
		timestamp: timestamp,
		isSell:    order.IsSell,
		price:     order.Price,
		volume:    order.Volume,
	}

	var tree, otherTree *btree.BTree
	if order.IsSell {
		tree = m.askTree
		otherTree = m.bidTree
	} else {
		tree = m.bidTree
		otherTree = m.askTree
	}

	switch order.Type {
	case pb.Order_MARKET:
		tradeEvents, orderBookEvents, err = m.processMarketOrder(tree, otherTree, item, order.Timestamp)
	case pb.Order_LIMIT:
		tradeEvents, orderBookEvents, err = m.processLimitOrder(tree, otherTree, item, order.Timestamp)
	case pb.Order_CANCEL:
		orderBookEvents, err = m.processCancelOrder(tree, item, order.Timestamp)
	default:
		err = errors.New("SubmitOrder(): unknown OrderType")
	}

	return
}

func (m *Matcher) processMarketOrder(tree, otherTree *btree.BTree, item orderItem,
	protoTimeNow *tspb.Timestamp) (tradeEvents []*pb.TradeEvent, orderBookEvents []*pb.OrderBookEvent, err error) {
	type volumeRecorder struct {
		volume     float64
		leftVolume float64
	}
	matchedMarketOrderMap := map[float64]volumeRecorder{}

	for item.volume > 0 && otherTree.Len() > 0 {
		maxItem := otherTree.Max().(orderItem)
		var matchedVolume float64

		if item.volume >= maxItem.volume {
			matchedVolume = maxItem.volume
		} else {
			matchedVolume = item.volume

			// Add residual volume of the limit order as a new limit order.
			residualMaxItem := maxItem
			residualMaxItem.volume -= matchedVolume
			ts, os, err2 := m.processLimitOrder(otherTree, tree, residualMaxItem, protoTimeNow)
			if err2 != nil {
				err = err2
			}
			tradeEvents = append(tradeEvents, ts...)
			orderBookEvents = append(orderBookEvents, os...)
		}

		// Cancel matched limit order.
		os, err2 := m.processCancelOrder(otherTree, maxItem, protoTimeNow)
		if err2 != nil {
			err = err2
		}
		orderBookEvents = append(orderBookEvents, os...)

		tradeEvents = append(tradeEvents, &pb.TradeEvent{
			OrderId:       maxItem.orderId,
			Timestamp:     protoTimeNow,
			IsTaker:       false,
			IsSell:        !item.isSell,
			Price:         maxItem.price,
			MatchedVolume: matchedVolume,
			LeftVolume:    maxItem.volume - matchedVolume,
		})

		item.volume -= matchedVolume
		accumulatedVolume := matchedVolume
		if _, ok := matchedMarketOrderMap[maxItem.price]; ok {
			accumulatedVolume = matchedVolume + matchedMarketOrderMap[maxItem.price].volume
		}
		matchedMarketOrderMap[maxItem.price] = volumeRecorder{
			volume:     accumulatedVolume,
			leftVolume: item.volume,
		}
	}

	// TODO: Send matched market order to notification channel.
	for price, volumeRecorder := range matchedMarketOrderMap {
		tradeEvents = append(tradeEvents, &pb.TradeEvent{
			OrderId:       item.orderId,
			Timestamp:     protoTimeNow,
			IsTaker:       true,
			IsSell:        item.isSell,
			Price:         price,
			MatchedVolume: volumeRecorder.volume,
			LeftVolume:    volumeRecorder.leftVolume,
		})
	}

	return
}

func (m *Matcher) processLimitOrder(tree, otherTree *btree.BTree, item orderItem,
	protoTimeNow *tspb.Timestamp) (tradeEvents []*pb.TradeEvent, orderBookEvents []*pb.OrderBookEvent, err error) {
	// Convert the limit order to market order if the order price is equal to or better than the best
	// price of the other tree.
	bestPrice := otherTree.Max().(orderItem).price
	if (item.price == bestPrice) || (item.isSell && item.price < bestPrice) {
		tradeEvents, orderBookEvents, err = m.processMarketOrder(tree, otherTree, item, protoTimeNow)
	}

	// Add the limit order to the order book.
	if tree.ReplaceOrInsert(item) != nil {
		err = errors.New("processLimitOrder(): limit order already exists")
	} else {
		orderBookEvents = append(orderBookEvents, &pb.OrderBookEvent{
			OrderId:   item.orderId,
			Timestamp: protoTimeNow,
			Type:      pb.OrderBookEvent_ADD,
			IsSell:    item.isSell,
			Price:     item.price,
			Volume:    item.volume,
		})
	}

	return
}

func (m *Matcher) processCancelOrder(tree *btree.BTree, item orderItem,
	protoTimeNow *tspb.Timestamp) (orderBookEvents []*pb.OrderBookEvent, err error) {
	if tree.Delete(item) == nil {
		err = errors.New("processCancelOrder(): order cannot be canceled: not exist")
	} else {
		orderBookEvents = append(orderBookEvents, &pb.OrderBookEvent{
			OrderId:   item.orderId,
			Timestamp: protoTimeNow,
			Type:      pb.OrderBookEvent_REMOVE,
			IsSell:    item.isSell,
			Price:     item.price,
			Volume:    item.volume,
		})
	}
	return
}
