package service

import (
	"testing"
	"time"

	mpb "github.com/catortiger/exchange/matching_engine/matcher/proto"
)

type fakeMatcher struct {
	tradingPair     mpb.TradingPair
	orderBook       *mpb.OrderBook
	tradeEvents     []*mpb.TradeEvent
	orderBookEvents []*mpb.OrderBookEvent
}

func (m *fakeMatcher) TradingPair() mpb.TradingPair {
	return m.tradingPair
}

func (m *fakeMatcher) OrderBook(snapshotTime time.Time) (*mpb.OrderBook, error) {
	return m.orderBook, nil
}

func (m *fakeMatcher) CreateOrder(order *mpb.Order) ([]*mpb.TradeEvent,
	[]*mpb.OrderBookEvent, error) {
	return m.tradeEvents, m.orderBookEvents, nil
}

func TestGetOrderBook(t *testing.T) {
}

func TestCreateOrder(t *testing.T) {
}
