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
	var tree, otherTree *btree.BTree
	var item btree.Item
	timestamp, err := ptypes.Timestamp(orderItem.Timestamp)
	if err != nil {
		return err
	}
	switch orderItem.Position {
	case pb.OrderItem_ASK:
		tree = askTree
		otherTree = bidTree
		item = askItem{
			orderId:   orderItem.OrderId,
			timestamp: timestamp,
			price:     orderItem.Price,
		}
	case pb.OrderItem_BID:
		tree = bidTree
		otherTree = askTree
		item = bidItem{
			orderId:   orderItem.OrderId,
			timestamp: timestamp,
			price:     orderItem.Price,
		}
	default:
		return errors.New("unknown OrderItem_Position")
	}

	switch orderItem.Type {
	case pb.OrderItem_MARKET:
		m.processMarketOrder(otherTree, item, orderItem.Position)
	case pb.OrderItem_LIMIT:
		m.processLimitOrder(tree, otherTree, item, orderItem.Position)
	case pb.OrderItem_CANCEL:
		m.processCancelOrder(tree, item)
	default:
		return errors.New("unknown OrderItem_Type")
	}

	return nil
}

func (m *Matcher) processMarketOrder(otherTree *btree.BTree, item btree.Item,
	position pb.OrderItem_Position) error {
	return nil
}

func (m *Matcher) processLimitOrder(tree, otherTree *btree.BTree, item btree.Item,
	position pb.OrderItem_Position) error {
	// Convert the limit order to market order if the order price is better than the best price of
	// the other tree.
	bestPriceOfOtherTree := otherTree.Max().Price
	if (pb.OrderItem_ASK && item.Price < bestPriceOfOtherTree) ||
		(pb.OrderItem_BID && item.Price > bestPriceOfOtherTree) {
		return m.processMarketOrder(otherTree, item, position)
	}

	// Add the limit order to the order book.
	if tree.ReplaceOrInsert(item) != nil {
		return errors.New("limit order already exists")
	} else {
		// TODO:Send limit order event to channel.
	}

	return nil
}

func (m *Matcher) processCancelOrder(tree *btree.BTree, item btree.Item) error {
	if tree.Delete(item) == nil {
		return errors.New("order cannot be canceled: not exist")
	} else {
		// TODO:Send cancel order event to channel.
	}
	return nil
}
