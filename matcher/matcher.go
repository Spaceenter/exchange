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
		isAsk:     order.OrderInfo.Position == pb.Position_ASK,
		price:     order.Price,
		volume:    order.Volume,
	}

	var tree, otherTree *btree.BTree
	switch order.OrderInfo.Position {
	case pb.Position_ASK:
		tree = m.askTree
		otherTree = m.bidTree
	case pb.Position_BID:
		tree = m.bidTree
		otherTree = m.askTree
	default:
		return errors.New("unknown Position")
	}

	switch order.OrderInfo.Type {
	case pb.OrderType_MARKET:
		return m.processMarketOrder(otherTree, item)
	case pb.OrderType_LIMIT:
		return m.processLimitOrder(tree, otherTree, item)
	case pb.OrderType_CANCEL:
		return m.processCancelOrder(tree, item)
	default:
		return errors.New("unknown OrderType")
	}

	return nil
}

func (m *Matcher) processMarketOrder(otherTree *btree.BTree, item orderItem) error {
	for item.volume > 0 && otherTree.Len() > 0 {
		//		maxItem := otherTree.Max().(orderItem)
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
