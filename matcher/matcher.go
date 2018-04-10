// Package matcher is an engine that receives and matches orders.
package matcher

import (
	"errors"

	"github.com/golang/protobuf/ptypes"
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

func (m *Matcher) SubmitOrder(orderItem pb.OrderItem) error {
	var tree *btree.BTree
	var item btree.Item
	timestamp, err := ptypes.Timestamp(orderItem.Timestamp)
	if err != nil {
		return err
	}
	switch orderItem.Position {
	case pb.OrderItem_ASK:
		tree = askTree
		item = askItem{
			orderId:   orderItem.OrderId,
			timestamp: timestamp,
			value:     orderItem.Value,
		}
	case pb.OrderItem_BID:
		tree = bidTree
		item = bidItem{
			orderId:   orderItem.OrderId,
			timestamp: timestamp,
			value:     orderItem.Value,
		}
	default:
		return errors.New("unknown OrderItem_Position")
	}

	switch orderItem.Type {
	case pb.OrderItem_MARKET:
		m.processMarketOrder(tree, item, orderItem.Position)
	case pb.OrderItem_LIMIT:
		m.processLimitOrder(tree, item, orderItem.Position)
	case pb.OrderItem_CANCEL:
		m.processCancelOrder(tree, item, orderItem.Position)
	default:
		return errors.New("unknown OrderItem_Type")
	}

	return nil
}

func (m *Matcher) processMarketOrder(tree *btree.BTree, item btree.Item,
	position pb.OrderItem_Position) error {
	return nil
}

func (m *Matcher) processLimitOrder(tree *btree.BTree, item btree.Item,
	position pb.OrderItem_Position) error {
	return nil
}

func (m *Matcher) processCancelOrder(tree *btree.BTree, item btree.Item,
	position pb.OrderItem_Position) error {
	return nil
}
