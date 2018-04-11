// Package matcher is an engine that receives and matches orders.
package matcher

import (
	"errors"
	"time"

	"github.com/golang/protobuf/ptypes"
	tspb "github.com/golang/protobuf/ptypes/timestamp"
	"github.com/google/btree"
	pb "github.com/spaceenter/exchange/proto"
)

type Matcher struct {
	tradingPair pb.TradingPair
	askTree     *btree.BTree
	bidTree     *btree.BTree
	// TODO: Add notification channel.
}

func New(tradingPair pb.TradingPair) *Matcher {
	return &Matcher{
		tradingPair: tradingPair,
		askTree:     btree.New(2),
		bidTree:     btree.New(2),
	}
}

func (m *Matcher) SubmitOrder(order pb.Order) error {
	timestamp, err := ptypes.Timestamp(order.OrderInfo.Timestamp)
	if err != nil {
		return err
	}
	item := orderItem{
		orderId:   order.OrderInfo.OrderId,
		timestamp: timestamp,
		isAsk:     order.OrderInfo.IsAsk,
		price:     order.Price,
		volume:    order.Volume,
	}

	var tree, otherTree *btree.BTree
	if order.OrderInfo.IsAsk {
		tree = m.askTree
		otherTree = m.bidTree
	} else {
		tree = m.bidTree
		otherTree = m.askTree
	}

	switch order.OrderInfo.Type {
	case pb.OrderType_MARKET:
		return m.processMarketOrder(otherTree, item)
	case pb.OrderType_LIMIT:
		return m.processLimitOrder(tree, otherTree, item)
	case pb.OrderType_CANCEL:
		return m.processCancelOrder(tree, item)
	default:
		return errors.New("SubmitOrder(): unknown OrderType")
	}

	return nil
}

func (m *Matcher) processMarketOrder(otherTree *btree.BTree, item orderItem) error {
	protoTimeNow, err := ptypes.TimestampProto(time.Now())
	if err != nil {
		return err
	}
	type volumeRecorder struct {
		volume     float64
		leftVolume float64
	}
	var matchedMarketOrderMap map[float64]volumeRecorder

	for item.volume > 0 && otherTree.Len() > 0 {
		maxItem := otherTree.Max().(orderItem)
		var matchedVolume float64

		if item.volume >= maxItem.volume {
			matchedVolume = maxItem.volume

			_ = matchedLimitOrder(maxItem, protoTimeNow, !item.isAsk, matchedVolume)
			if otherTree.Delete(maxItem) == nil {
				return errors.New("processMarketOrder(): cannot delete limit order")
			}
			// TODO: send the matchedLimitOrder to notification channel.
		} else { //item.volume < maxItem.volume
			matchedVolume = item.volume

			_ = matchedLimitOrder(maxItem, protoTimeNow, !item.isAsk, matchedVolume)
			if otherTree.Delete(maxItem) == nil {
				return errors.New("processMarketOrder(): cannot delete limit order")
			}
			// TODO: send the matchedLimitOrder to notification channel.
			maxItem.volume -= matchedVolume
			if otherTree.ReplaceOrInsert(maxItem) != nil {
				return errors.New("processMarketOrder(): limit order already exists")
			} else {
				// TODO: Send limit order event to channel.
			}
		}

		item.volume -= matchedVolume
		if _, ok := matchedMarketOrderMap[maxItem.price]; ok {
			vr := matchedMarketOrderMap[maxItem.price]
			vr.volume += matchedVolume
			vr.leftVolume = item.volume
			matchedMarketOrderMap[maxItem.price] = vr
		} else {
			matchedMarketOrderMap[maxItem.price] = volumeRecorder{
				volume:     matchedVolume,
				leftVolume: item.volume,
			}
		}
	}

	// TODO: Send matched market order to notification channel.
	for price, volumeRecorder := range matchedMarketOrderMap {
		_ = &pb.MatchedOrder{
			OrderInfo: &pb.OrderInfo{
				OrderId:   item.orderId,
				Timestamp: protoTimeNow,
				Type:      pb.OrderType_MARKET,
				IsAsk:     item.isAsk,
			},
			MatchedPrice:  price,
			MatchedVolume: volumeRecorder.volume,
			LeftVolume:    volumeRecorder.leftVolume,
		}
	}

	return nil
}

func (m *Matcher) processLimitOrder(tree, otherTree *btree.BTree, item orderItem) error {
	// Convert the limit order to market order if the order price is equal to or better than the best
	// price of the other tree.
	bestPrice := otherTree.Max().(orderItem).price
	if (item.price == bestPrice) || (item.isAsk && item.price < bestPrice) {
		return m.processMarketOrder(otherTree, item)
	}

	// Add the limit order to the order book.
	if tree.ReplaceOrInsert(item) != nil {
		return errors.New("processLimitOrder(): limit order already exists")
	} else {
		// TODO: Send limit order event to channel.
	}

	return nil
}

func (m *Matcher) processCancelOrder(tree *btree.BTree, item btree.Item) error {
	if tree.Delete(item) == nil {
		return errors.New("processCancelOrder(): order cannot be canceled: not exist")
	} else {
		// TODO: Send cancel order event to channel.
	}
	return nil
}

func matchedLimitOrder(item orderItem, timestamp *tspb.Timestamp, isAsk bool,
	matchedVolume float64) *pb.MatchedOrder {
	return &pb.MatchedOrder{
		OrderInfo: &pb.OrderInfo{
			OrderId:   item.orderId,
			Timestamp: timestamp,
			Type:      pb.OrderType_LIMIT,
			IsAsk:     isAsk,
		},
		MatchedPrice:  item.price,
		MatchedVolume: matchedVolume,
		LeftVolume:    item.volume - matchedVolume,
	}
}
