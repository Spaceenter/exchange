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
	timestamp, err := ptypes.Timestamp(order.Timestamp)
	if err != nil {
		return err
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
		return m.processMarketOrder(otherTree, item)
	case pb.Order_LIMIT:
		return m.processLimitOrder(tree, otherTree, item)
	case pb.Order_CANCEL:
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
	matchedMarketOrderMap := map[float64]volumeRecorder{}

	for item.volume > 0 && otherTree.Len() > 0 {
		maxItem := otherTree.Max().(orderItem)
		var matchedVolume float64

		if item.volume >= maxItem.volume {
			matchedVolume = maxItem.volume

			if otherTree.Delete(maxItem) == nil {
				return errors.New("processMarketOrder(): cannot delete limit order")
			}
			// TODO: send the matchedLimitOrder to notification channel.
		} else {
			matchedVolume = item.volume

			_ = tradeEvent(maxItem, protoTimeNow, !item.isSell, false, matchedVolume)
			// TODO: send the matchedLimitOrder to notification channel.
			if otherTree.Delete(maxItem) == nil {
				return errors.New("processMarketOrder(): cannot delete limit order")
			}
			// TODO: send cancel order (followed by adding remaining balance) to notification channel.
			maxItem.volume -= matchedVolume
			if otherTree.ReplaceOrInsert(maxItem) != nil {
				return errors.New("processMarketOrder(): limit order already exists")
			} else {
				// TODO: Send limit order event to channel.
			}
		}

		_ = &pb.TradeEvent{
			OrderId:       maxItem.orderId,
			Timestamp:     protoTimeNow,
			IsTaker:       false,
			IsSell:        !item.isSell,
			Price:         maxItem.price,
			MatchedVolume: matchedVolume,
			LeftVolume:    maxItem.volume - matchedVolume,
		}

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
		_ = &pb.TradeEvent{
			OrderId:       item.orderId,
			Timestamp:     protoTimeNow,
			IsTaker:       true,
			IsSell:        item.isSell,
			Price:         price,
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
	if (item.price == bestPrice) || (item.isSell && item.price < bestPrice) {
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

func tradeEvent(item orderItem, timestamp *tspb.Timestamp, isSell, isTaker bool,
	matchedVolume float64) *pb.TradeEvent {
	return &pb.TradeEvent{
		OrderId:       item.orderId,
		Timestamp:     timestamp,
		IsTaker:       isTaker,
		IsSell:        isSell,
		Price:         item.price,
		MatchedVolume: matchedVolume,
		LeftVolume:    item.volume - matchedVolume,
	}
}
