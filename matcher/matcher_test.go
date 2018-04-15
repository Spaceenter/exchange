package matcher

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	pb "github.com/spaceenter/exchange/proto"
	"github.com/spaceenter/exchange/testutil"
)

func TestSimpleLimitOrderAndCancelOrder(t *testing.T) {
	matcher := New(pb.TradingPair_BTC_USD)

	// Add simple limit orders.
	for _, c := range []struct {
		order               *pb.Order
		wantOrderBookEvents []*pb.OrderBookEvent
	}{
		{
			&pb.Order{
				OrderId:   "a",
				OrderTime: testutil.ProtoTimeEarly,
				Type:      pb.Order_LIMIT,
				IsSell:    true,
				Price:     2.1,
				Volume:    3.2,
			},
			[]*pb.OrderBookEvent{
				{
					OrderId:   "a",
					Timestamp: testutil.ProtoTimeEarly,
					Type:      pb.OrderBookEvent_ADD,
					IsSell:    true,
					Price:     2.1,
					Volume:    3.2,
				},
			},
		},
		{
			&pb.Order{
				OrderId:   "b",
				OrderTime: testutil.ProtoTimeNow,
				Type:      pb.Order_LIMIT,
				IsSell:    false,
				Price:     2.2,
				Volume:    3.3,
			},
			[]*pb.OrderBookEvent{
				{
					OrderId:   "b",
					Timestamp: testutil.ProtoTimeNow,
					Type:      pb.OrderBookEvent_ADD,
					IsSell:    false,
					Price:     2.2,
					Volume:    3.3,
				},
			},
		},
	} {
		gotTradeEvents, gotOrderBookEvents, err := matcher.SubmitOrder(c.order)
		if err != nil {
			t.Errorf("SubmitOrder(LIMIT) = %s", err)
			continue
		}
		if gotTradeEvents != nil {
			t.Errorf("SubmitOrder(LIMIT) TradeEvents = %s, want nil", gotTradeEvents)
		}
		if !cmp.Equal(gotOrderBookEvents, c.wantOrderBookEvents) {
			t.Errorf("SubmitOrder(LIMIT) OrderBookEvents = %s, want %s", gotOrderBookEvents,
				c.wantOrderBookEvents)
		}
	}

	// Check order book after adding simple limit orders.
	gotOrderBook, err := matcher.OrderBook(testutil.TimeLate)
	if err != nil {
		t.Fatalf("OrderBook(LIMIT) = %s", err)
	}
	wantOrderBook := &pb.OrderBook{
		Pair:         pb.TradingPair_BTC_USD,
		SnapshotTime: testutil.ProtoTimeLate,
		Trees: []*pb.OrderTree{
			{
				IsSell: true,
				Items: []*pb.OrderItem{
					{
						OrderId:   "a",
						OrderTime: testutil.ProtoTimeEarly,
						Price:     2.1,
						Volume:    3.2,
					},
				},
			},
			{
				IsSell: false,
				Items: []*pb.OrderItem{
					{
						OrderId:   "b",
						OrderTime: testutil.ProtoTimeNow,
						Price:     2.2,
						Volume:    3.3,
					},
				},
			},
		},
	}
	if !cmp.Equal(gotOrderBook, wantOrderBook) {
		t.Errorf("OrderBook(LIMIT) = %s, want %s", gotOrderBook, wantOrderBook)
	}

	// Cancel order.
	gotTradeEvents, gotOrderBookEvents, err := matcher.SubmitOrder(&pb.Order{
		OrderId:   "a",
		OrderTime: testutil.ProtoTimeEarly,
		Type:      pb.Order_CANCEL,
		IsSell:    true,
		Price:     2.1,
		Volume:    3.2,
	})
	if gotTradeEvents != nil {
		t.Errorf("SubmitOrder(CANCEL) TradeEvents = %s, want nil", gotTradeEvents)
	}
	wantOrderBookEvents := []*pb.OrderBookEvent{
		{
			OrderId:   "a",
			Timestamp: testutil.ProtoTimeEarly,
			Type:      pb.OrderBookEvent_REMOVE,
			IsSell:    true,
			Price:     2.1,
			Volume:    3.2,
		},
	}
	if !cmp.Equal(gotOrderBookEvents, wantOrderBookEvents) {
		t.Errorf("SubmitOrder(CANCEL) OrderBookEvents = %s, want %s", gotOrderBookEvents,
			wantOrderBookEvents)
	}

	// Check order book after cancel order.
	gotOrderBook, err = matcher.OrderBook(testutil.TimeLate)
	if err != nil {
		t.Fatalf("OrderBook(CANCEL) = %s", err)
	}
	wantOrderBook = &pb.OrderBook{
		Pair:         pb.TradingPair_BTC_USD,
		SnapshotTime: testutil.ProtoTimeLate,
		Trees: []*pb.OrderTree{
			{
				IsSell: false,
				Items: []*pb.OrderItem{
					{
						OrderId:   "b",
						OrderTime: testutil.ProtoTimeNow,
						Price:     2.2,
						Volume:    3.3,
					},
				},
			},
		},
	}
	if !cmp.Equal(gotOrderBook, wantOrderBook) {
		t.Errorf("OrderBook(CANCEL) = %s, want %s", gotOrderBook, wantOrderBook)
	}
}

func TestLimitOrderMultiple(t *testing.T) {
	matcher := New(pb.TradingPair_BTC_USD)

	// Add limit orders as makers.
	for _, l := range []struct {
		orderId string
		isSell  bool
		price   float64
		volume  float64
	}{
		{"s1", true, 2.1, 2},
		{"s2", true, 2.1, 2},
		{"s3", true, 2.2, 2},
		{"s4", true, 2.3, 3},
		{"b1", true, 1.9, 2},
		{"b2", true, 1.9, 2},
		{"b3", true, 1.8, 2},
		{"b4", true, 1.7, 3},
	} {
		order := &pb.Order{
			OrderId:   l.orderId,
			OrderTime: testutil.ProtoTimeEarly,
			Type:      pb.Order_LIMIT,
			IsSell:    l.isSell,
			Price:     l.price,
			Volume:    l.volume,
		}
		if _, _, err := matcher.SubmitOrder(order); err != nil {
			t.Errorf("SubmitOrder(%s) = %s", l.orderId, err)
		}
	}
}

func TestMarketOrder(t *testing.T) {
}

func TestLimitOrderConvertToMarketOrder(t *testing.T) {
	// matcher := New(pb.TradingPair_BTC_USD)
}
