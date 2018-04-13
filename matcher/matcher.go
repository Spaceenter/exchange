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

// OrderBook gets the orderbook.
// This function should be called infrequently to ensure the performance of the matching engine.
// A separate routine should be used for updating the orderbook by listening to OrderBookEvent.
// This function should only be served as a periodical check point to ensure the correctness.
func (m *Matcher) OrderBook(snapshotTime time.Time) (*pb.OrderBook, error) {
	snapshotTimeProto, err := ptypes.TimestampProto(snapshotTime)
	if err != nil {
		return nil, err
	}
	orderBook := &pb.OrderBook{
		Pair:         m.tradingPair,
		SnapshotTime: snapshotTimeProto,
	}
	for _, t := range []struct {
		isSell bool
		tree   *btree.BTree
	}{
		{true, m.sellTree},
		{false, m.buyTree},
	} {
		orderTree := &pb.OrderTree{IsSell: t.isSell}
		t.tree.Descend(btree.ItemIterator(func(i btree.Item) bool {
			item := i.(orderItem)
			orderTimeProto, err := ptypes.TimestampProto(item.orderTime)
			if err != nil {
				return false
			}
			orderTree.Items = append(orderTree.Items, &pb.OrderItem{
				OrderId:   item.orderId,
				OrderTime: orderTimeProto,
				Price:     item.price,
				Volume:    item.volume,
			})
			return true
		}))
		orderBook.Trees = append(orderBook.Trees, orderTree)
	}
	return orderBook, nil
}

// SubmitOrder submits an order, and gets corresponding trade and order book events.
func (m *Matcher) SubmitOrder(order *pb.Order) ([]*pb.TradeEvent, []*pb.OrderBookEvent, error) {
	orderTime, err := ptypes.Timestamp(order.OrderTime)
	if err != nil {
		return nil, nil, err
	}
	item := orderItem{
		orderId:   order.OrderId,
		orderTime: orderTime,
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
		return m.processMarketOrder(tree, otherTree, item, order.OrderTime)
	case pb.Order_LIMIT:
		return m.processLimitOrder(tree, otherTree, item, order.OrderTime)
	case pb.Order_CANCEL:
		return m.processCancelOrder(tree, item, order.OrderTime)
	default:
		return nil, nil, errors.New("SubmitOrder(): unknown OrderType")
	}
}

func (m *Matcher) processMarketOrder(tree, otherTree *btree.BTree, item orderItem,
	orderTimeProto *tspb.Timestamp) ([]*pb.TradeEvent, []*pb.OrderBookEvent, error) {
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
			ts, os, err := m.processLimitOrder(otherTree, tree, residualMaxItem, orderTimeProto)
			if err != nil {
				return nil, nil, err
			}
			tradeEvents = append(tradeEvents, ts...)
			orderBookEvents = append(orderBookEvents, os...)
		}

		// Cancel matched limit order.
		_, os, err := m.processCancelOrder(otherTree, maxItem, orderTimeProto)
		if err != nil {
			return nil, nil, err
		}
		orderBookEvents = append(orderBookEvents, os...)

		// Trade event of the matched limit order.
		tradeEvents = append(tradeEvents, &pb.TradeEvent{
			OrderId:       maxItem.orderId,
			Timestamp:     orderTimeProto,
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
			Timestamp:     orderTimeProto,
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
	orderTimeProto *tspb.Timestamp) ([]*pb.TradeEvent, []*pb.OrderBookEvent, error) {
	// Convert the limit order to market order if the order price is equal to or better than the best
	// price of the other tree.
	bestPrice := otherTree.Max().(orderItem).price
	if (item.price == bestPrice) || (item.isSell && item.price < bestPrice) {
		return m.processMarketOrder(tree, otherTree, item, orderTimeProto)
	}

	// Add the limit order to the order book.
	orderBookEvents := []*pb.OrderBookEvent{}
	if tree.ReplaceOrInsert(item) != nil {
		return nil, nil, errors.New("processLimitOrder(): limit order already exists")
	} else {
		orderBookEvents = append(orderBookEvents, &pb.OrderBookEvent{
			OrderId:   item.orderId,
			Timestamp: orderTimeProto,
			Type:      pb.OrderBookEvent_ADD,
			IsSell:    item.isSell,
			Price:     item.price,
			Volume:    item.volume,
		})
	}

	return nil, orderBookEvents, nil
}

func (m *Matcher) processCancelOrder(tree *btree.BTree, item orderItem,
	orderTimeProto *tspb.Timestamp) ([]*pb.TradeEvent, []*pb.OrderBookEvent, error) {
	orderBookEvents := []*pb.OrderBookEvent{}
	if tree.Delete(item) == nil {
		return nil, nil, errors.New("processCancelOrder(): order cannot be canceled: not exist")
	} else {
		orderBookEvents = append(orderBookEvents, &pb.OrderBookEvent{
			OrderId:   item.orderId,
			Timestamp: orderTimeProto,
			Type:      pb.OrderBookEvent_REMOVE,
			IsSell:    item.isSell,
			Price:     item.price,
			Volume:    item.volume,
		})
	}
	return nil, orderBookEvents, nil
}
