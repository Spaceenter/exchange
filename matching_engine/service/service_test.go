package service

import (
	"testing"
	"time"
)

type fakeStore struct{}

// TODO: Implement fakeStore functions to satisfy its interface.

type fakeMatcher struct {
	orderBook       *pb.OrderBook
	tradeEvents     []*pb.TradeEvent
	orderBookEvents []*pb.OrderBookEvent
}

func (m *fakeMatcher) OrderBook(snapshotTime time.Time) (*pb.OrderBook, error) {
	return m.orderBook, nil
}

func (m *fakeMatcher) CreateOrder(order *pb.Order) ([]*pb.TradeEvent, []*pb.OrderBookEvent, error) {
	return m.tradeEvents, m.orderBookEvents, nil
}

func TestGetOrderBook(t *testing.T) {
}

func TestCreateOrder(t *testing.T) {
}
