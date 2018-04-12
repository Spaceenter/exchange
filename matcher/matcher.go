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
	sellTree    *btree.BTree
	buyTree     *btree.BTree
}

func New(tradingPair pb.TradingPair) *Matcher {
	return &Matcher{
		tradingPair: tradingPair,
		sellTree:    btree.New(2),
		buyTree:     btree.New(2),
	}
}

func (m *Matcher) OrderBook() {
}

// SubmitOrder submits an order, and gets corresponding trade and order book events.
func (m *Matcher) SubmitOrder(order pb.Order) ([]*pb.TradeEvent, []*pb.OrderBookEvent, error) {
	timestamp, err := ptypes.Timestamp(order.Timestamp)
	if err != nil {
		return nil, nil, err
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
		tree = m.sellTree
		otherTree = m.buyTree
	} else {
		tree = m.buyTree
		otherTree = m.sellTree
	}

	switch order.Type {
	case pb.Order_MARKET:
		return m.processMarketOrder(tree, otherTree, item, order.Timestamp)
	case pb.Order_LIMIT:
		return m.processLimitOrder(tree, otherTree, item, order.Timestamp)
	case pb.Order_CANCEL:
		return m.processCancelOrder(tree, item, order.Timestamp)
	default:
		return nil, nil, errors.New("SubmitOrder(): unknown OrderType")
	}
}

func (m *Matcher) processMarketOrder(tree, otherTree *btree.BTree, item orderItem,
	timeNowProto *tspb.Timestamp) ([]*pb.TradeEvent, []*pb.OrderBookEvent, error) {
	tradeEvents := []*pb.TradeEvent{}
	orderBookEvents := []*pb.OrderBookEvent{}
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
			ts, os, err := m.processLimitOrder(otherTree, tree, residualMaxItem, timeNowProto)
			if err != nil {
				return nil, nil, err
			}
			tradeEvents = append(tradeEvents, ts...)
			orderBookEvents = append(orderBookEvents, os...)
		}

		// Cancel matched limit order.
		_, os, err := m.processCancelOrder(otherTree, maxItem, timeNowProto)
		if err != nil {
			return nil, nil, err
		}
		orderBookEvents = append(orderBookEvents, os...)

		// Trade event of the matched limit order.
		tradeEvents = append(tradeEvents, &pb.TradeEvent{
			OrderId:       maxItem.orderId,
			Timestamp:     timeNowProto,
			IsTaker:       false,
			IsSell:        !item.isSell,
			Price:         maxItem.price,
			MatchedVolume: matchedVolume,
			LeftVolume:    maxItem.volume - matchedVolume,
		})

		// Aggregate filled market order price and volume information.
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

	// Trade events of the matched market order.
	for price, volumeRecorder := range matchedMarketOrderMap {
		tradeEvents = append(tradeEvents, &pb.TradeEvent{
			OrderId:       item.orderId,
			Timestamp:     timeNowProto,
			IsTaker:       true,
			IsSell:        item.isSell,
			Price:         price,
			MatchedVolume: volumeRecorder.volume,
			LeftVolume:    volumeRecorder.leftVolume,
		})
	}

	return tradeEvents, orderBookEvents, nil
}

func (m *Matcher) processLimitOrder(tree, otherTree *btree.BTree, item orderItem,
	timeNowProto *tspb.Timestamp) ([]*pb.TradeEvent, []*pb.OrderBookEvent, error) {
	// Convert the limit order to market order if the order price is equal to or better than the best
	// price of the other tree.
	bestPrice := otherTree.Max().(orderItem).price
	if (item.price == bestPrice) || (item.isSell && item.price < bestPrice) {
		return m.processMarketOrder(tree, otherTree, item, timeNowProto)
	}

	// Add the limit order to the order book.
	orderBookEvents := []*pb.OrderBookEvent{}
	if tree.ReplaceOrInsert(item) != nil {
		return nil, nil, errors.New("processLimitOrder(): limit order already exists")
	} else {
		orderBookEvents = append(orderBookEvents, &pb.OrderBookEvent{
			OrderId:   item.orderId,
			Timestamp: timeNowProto,
			Type:      pb.OrderBookEvent_ADD,
			IsSell:    item.isSell,
			Price:     item.price,
			Volume:    item.volume,
		})
	}

	return nil, orderBookEvents, nil
}

func (m *Matcher) processCancelOrder(tree *btree.BTree, item orderItem,
	timeNowProto *tspb.Timestamp) ([]*pb.TradeEvent, []*pb.OrderBookEvent, error) {
	orderBookEvents := []*pb.OrderBookEvent{}
	if tree.Delete(item) == nil {
		return nil, nil, errors.New("processCancelOrder(): order cannot be canceled: not exist")
	} else {
		orderBookEvents = append(orderBookEvents, &pb.OrderBookEvent{
			OrderId:   item.orderId,
			Timestamp: timeNowProto,
			Type:      pb.OrderBookEvent_REMOVE,
			IsSell:    item.isSell,
			Price:     item.price,
			Volume:    item.volume,
		})
	}
	return nil, orderBookEvents, nil
}
