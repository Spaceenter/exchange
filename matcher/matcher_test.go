package matcher

import (
	"reflect"
	"testing"

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
		if !reflect.DeepEqual(gotOrderBookEvents, c.wantOrderBookEvents) {
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
	if !reflect.DeepEqual(gotOrderBook, wantOrderBook) {
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
	if !reflect.DeepEqual(gotOrderBookEvents, wantOrderBookEvents) {
		t.Errorf("SubmitOrder(CANCEL) OrderBookEvents = %s, want %s", gotOrderBookEvents,
			wantOrderBookEvents)
	}

	// Check order book.
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
	if !reflect.DeepEqual(gotOrderBook, wantOrderBook) {
		t.Errorf("OrderBook(CANCEL) = %s, want %s", gotOrderBook, wantOrderBook)
	}
}

func TestMarketOrder(t *testing.T) {
	matcher, err := buildTestOrderBook()
	if err != nil {
		t.Fatalf("buildTestOrderBook() = %s", err)
	}

	// Add market orders.
	for _, c := range []struct {
		order               *pb.Order
		wantTradeEvents     []*pb.TradeEvent
		wantOrderBookEvents []*pb.OrderBookEvent
	}{
		// A small market order.
		{
			&pb.Order{
				OrderId:   "a",
				OrderTime: testutil.ProtoTimeNow,
				Type:      pb.Order_MARKET,
				IsSell:    true,
				Volume:    0.5,
			},
			[]*pb.TradeEvent{
				{
					OrderId:       "b2",
					Timestamp:     testutil.ProtoTimeNow,
					IsTaker:       false,
					IsSell:        false,
					Price:         1.9,
					MatchedVolume: 0.5,
					LeftVolume:    1.5,
				},
				{
					OrderId:       "a",
					Timestamp:     testutil.ProtoTimeNow,
					IsTaker:       true,
					IsSell:        true,
					Price:         1.9,
					MatchedVolume: 0.5,
					LeftVolume:    0,
				},
			},
			[]*pb.OrderBookEvent{
				{
					OrderId:   "b2",
					Timestamp: testutil.ProtoTimeNow,
					Type:      pb.OrderBookEvent_REMOVE,
					IsSell:    false,
					Price:     1.9,
					Volume:    2,
				},
				{
					OrderId:   "b2",
					Timestamp: testutil.ProtoTimeNow,
					Type:      pb.OrderBookEvent_ADD,
					IsSell:    false,
					Price:     1.9,
					Volume:    1.5,
				},
			},
		},
		// a large market order, but not large enough to eat all the order book.
		{
			&pb.Order{
				OrderId:   "b",
				OrderTime: testutil.ProtoTimeNow,
				Type:      pb.Order_MARKET,
				IsSell:    false,
				Volume:    5,
			},
			[]*pb.TradeEvent{
				{
					OrderId:       "s2",
					Timestamp:     testutil.ProtoTimeNow,
					IsTaker:       false,
					IsSell:        true,
					Price:         2.1,
					MatchedVolume: 2,
					LeftVolume:    0,
				},
				{
					OrderId:       "s1",
					Timestamp:     testutil.ProtoTimeNow,
					IsTaker:       false,
					IsSell:        true,
					Price:         2.1,
					MatchedVolume: 2,
					LeftVolume:    0,
				},
				{
					OrderId:       "s3",
					Timestamp:     testutil.ProtoTimeNow,
					IsTaker:       false,
					IsSell:        true,
					Price:         2.2,
					MatchedVolume: 1,
					LeftVolume:    1,
				},
				{
					OrderId:       "b",
					Timestamp:     testutil.ProtoTimeNow,
					IsTaker:       true,
					IsSell:        false,
					Price:         2.1,
					MatchedVolume: 4,
					LeftVolume:    1,
				},
				{
					OrderId:       "b",
					Timestamp:     testutil.ProtoTimeNow,
					IsTaker:       true,
					IsSell:        false,
					Price:         2.2,
					MatchedVolume: 1,
					LeftVolume:    0,
				},
			},
			[]*pb.OrderBookEvent{
				{
					OrderId:   "s2",
					Timestamp: testutil.ProtoTimeNow,
					Type:      pb.OrderBookEvent_REMOVE,
					IsSell:    true,
					Price:     2.1,
					Volume:    2,
				},
				{
					OrderId:   "s1",
					Timestamp: testutil.ProtoTimeNow,
					Type:      pb.OrderBookEvent_REMOVE,
					IsSell:    true,
					Price:     2.1,
					Volume:    2,
				},
				{
					OrderId:   "s3",
					Timestamp: testutil.ProtoTimeNow,
					Type:      pb.OrderBookEvent_REMOVE,
					IsSell:    true,
					Price:     2.2,
					Volume:    2,
				},
				{
					OrderId:   "s3",
					Timestamp: testutil.ProtoTimeNow,
					Type:      pb.OrderBookEvent_ADD,
					IsSell:    true,
					Price:     2.2,
					Volume:    1,
				},
			},
		},
		// a market order that is enough to eat all the rest order book.
		{
			&pb.Order{
				OrderId:   "c",
				OrderTime: testutil.ProtoTimeNow,
				Type:      pb.Order_MARKET,
				IsSell:    false,
				Volume:    3,
			},
			[]*pb.TradeEvent{
				{
					OrderId:       "s3",
					Timestamp:     testutil.ProtoTimeNow,
					IsTaker:       false,
					IsSell:        true,
					Price:         2.2,
					MatchedVolume: 1,
					LeftVolume:    0,
				},
				{
					OrderId:       "c",
					Timestamp:     testutil.ProtoTimeNow,
					IsTaker:       true,
					IsSell:        false,
					Price:         2.2,
					MatchedVolume: 1,
					LeftVolume:    2,
				},
			},
			[]*pb.OrderBookEvent{
				{
					OrderId:   "s3",
					Timestamp: testutil.ProtoTimeNow,
					Type:      pb.OrderBookEvent_REMOVE,
					IsSell:    true,
					Price:     2.2,
					Volume:    1,
				},
			},
		},
	} {
		gotTradeEvents, gotOrderBookEvents, err := matcher.SubmitOrder(c.order)
		if err != nil {
			t.Errorf("SubmitOrder(%s) = %s", c.order.OrderId, err)
			continue
		}
		if !reflect.DeepEqual(gotTradeEvents, c.wantTradeEvents) {
			t.Errorf("SubmitOrder(%s) TradeEvents = %s, want %s", c.order.OrderId, gotTradeEvents,
				c.wantTradeEvents)
		}
		if !reflect.DeepEqual(gotOrderBookEvents, c.wantOrderBookEvents) {
			t.Errorf("SubmitOrder(%s) OrderBookEvents = %s, want %s", c.order.OrderId, gotOrderBookEvents,
				c.wantOrderBookEvents)
		}
	}

	// Check order book.
	gotOrderBook, err := matcher.OrderBook(testutil.TimeLate)
	if err != nil {
		t.Fatalf("OrderBook() = %s", err)
	}
	wantOrderBook := &pb.OrderBook{
		Pair:         pb.TradingPair_BTC_USD,
		SnapshotTime: testutil.ProtoTimeLate,
		Trees: []*pb.OrderTree{
			{
				IsSell: false,
				Items: []*pb.OrderItem{
					{
						OrderId:   "b2",
						OrderTime: testutil.ProtoTimeEarly,
						Price:     1.9,
						Volume:    1.5,
					},
					{
						OrderId:   "b1",
						OrderTime: testutil.ProtoTimeEarly,
						Price:     1.9,
						Volume:    2,
					},
					{
						OrderId:   "b3",
						OrderTime: testutil.ProtoTimeEarly,
						Price:     1.8,
						Volume:    2,
					},
				},
			},
		},
	}
	if !reflect.DeepEqual(gotOrderBook, wantOrderBook) {
		t.Errorf("OrderBook() = %s, want %s", gotOrderBook, wantOrderBook)
	}
}

func TestLimitOrderConvertToMarketOrder(t *testing.T) {
	/*
		matcher, err := buildTestOrderBook()
		if err != nil {
			t.Fatalf("buildTestOrderBook() = %s", err)
		}
	*/
}

func buildTestOrderBook() (*Matcher, error) {
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
		{"b1", false, 1.9, 2},
		{"b2", false, 1.9, 2},
		{"b3", false, 1.8, 2},
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
			return matcher, err
		}
	}

	return matcher, nil
}
