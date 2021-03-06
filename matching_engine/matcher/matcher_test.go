package matcher

import (
	"reflect"
	"testing"

	pb "github.com/catortiger/exchange/matching_engine/matcher/proto"
	"github.com/catortiger/exchange/testutil"
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
				Price:     1.9,
				Volume:    3.3,
			},
			[]*pb.OrderBookEvent{
				{
					OrderId:   "b",
					Timestamp: testutil.ProtoTimeNow,
					Type:      pb.OrderBookEvent_ADD,
					IsSell:    false,
					Price:     1.9,
					Volume:    3.3,
				},
			},
		},
	} {
		gotTradeEvents, gotOrderBookEvents, err := matcher.CreateOrder(c.order)
		if err != nil {
			t.Errorf("CreateOrder(LIMIT) = %s", err)
			continue
		}
		if gotTradeEvents != nil {
			t.Errorf("CreateOrder(LIMIT) TradeEvents = %s, want nil", gotTradeEvents)
		}
		if !reflect.DeepEqual(gotOrderBookEvents, c.wantOrderBookEvents) {
			t.Errorf("CreateOrder(LIMIT) OrderBookEvents = %s, want %s", gotOrderBookEvents,
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
						Price:     1.9,
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
	gotTradeEvents, gotOrderBookEvents, err := matcher.CreateOrder(&pb.Order{
		OrderId:   "a",
		OrderTime: testutil.ProtoTimeEarly,
		Type:      pb.Order_CANCEL,
		IsSell:    true,
		Price:     2.1,
		Volume:    3.2,
	})
	if gotTradeEvents != nil {
		t.Errorf("CreateOrder(CANCEL) TradeEvents = %s, want nil", gotTradeEvents)
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
		t.Errorf("CreateOrder(CANCEL) OrderBookEvents = %s, want %s", gotOrderBookEvents,
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
						Price:     1.9,
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
		gotTradeEvents, gotOrderBookEvents, err := matcher.CreateOrder(c.order)
		if err != nil {
			t.Errorf("CreateOrder(%s) = %s", c.order.OrderId, err)
			continue
		}
		gotTradeEventsMap := tradeEventsSliceToMap(gotTradeEvents)
		wantTradeEventsMap := tradeEventsSliceToMap(c.wantTradeEvents)
		if !reflect.DeepEqual(gotTradeEventsMap, wantTradeEventsMap) {
			t.Errorf("CreateOrder(%s) TradeEvents = %v, want %v", c.order.OrderId, gotTradeEventsMap,
				wantTradeEventsMap)
		}
		if !reflect.DeepEqual(gotOrderBookEvents, c.wantOrderBookEvents) {
			t.Errorf("CreateOrder(%s) OrderBookEvents = %s, want %s", c.order.OrderId, gotOrderBookEvents,
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
	matcher, err := buildTestOrderBook()
	if err != nil {
		t.Fatalf("buildTestOrderBook() = %s", err)
	}

	for _, c := range []struct {
		order               *pb.Order
		wantTradeEvents     []*pb.TradeEvent
		wantOrderBookEvents []*pb.OrderBookEvent
	}{
		// Small limit order with price far away from middle price.
		{
			&pb.Order{
				OrderId:   "a",
				OrderTime: testutil.ProtoTimeNow,
				Type:      pb.Order_LIMIT,
				IsSell:    true,
				Price:     1.7,
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
		// Large limit order with price close to middle price.
		{
			&pb.Order{
				OrderId:   "b",
				OrderTime: testutil.ProtoTimeNow,
				Type:      pb.Order_LIMIT,
				IsSell:    false,
				Price:     2.15,
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
					OrderId:       "b",
					Timestamp:     testutil.ProtoTimeNow,
					IsTaker:       true,
					IsSell:        false,
					Price:         2.1,
					MatchedVolume: 4,
					LeftVolume:    1,
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
					OrderId:   "b",
					Timestamp: testutil.ProtoTimeNow,
					Type:      pb.OrderBookEvent_ADD,
					IsSell:    false,
					Price:     2.15,
					Volume:    1,
				},
			},
		},
	} {
		gotTradeEvents, gotOrderBookEvents, err := matcher.CreateOrder(c.order)
		if err != nil {
			t.Errorf("CreateOrder(%s) = %s", c.order.OrderId, err)
			continue
		}
		gotTradeEventsMap := tradeEventsSliceToMap(gotTradeEvents)
		wantTradeEventsMap := tradeEventsSliceToMap(c.wantTradeEvents)
		if !reflect.DeepEqual(gotTradeEventsMap, wantTradeEventsMap) {
			t.Errorf("CreateOrder(%s) TradeEvents = %v, want %v", c.order.OrderId, gotTradeEventsMap,
				wantTradeEventsMap)
		}
		if !reflect.DeepEqual(gotOrderBookEvents, c.wantOrderBookEvents) {
			t.Errorf("CreateOrder(%s) OrderBookEvents = %s, want %s", c.order.OrderId, gotOrderBookEvents,
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
				IsSell: true,
				Items: []*pb.OrderItem{
					{
						OrderId:   "s3",
						OrderTime: testutil.ProtoTimeEarly,
						Price:     2.2,
						Volume:    2,
					},
				},
			},
			{
				IsSell: false,
				Items: []*pb.OrderItem{
					{
						OrderId:   "b",
						OrderTime: testutil.ProtoTimeNow,
						Price:     2.15,
						Volume:    1,
					},
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
		if _, _, err := matcher.CreateOrder(order); err != nil {
			return matcher, err
		}
	}

	return matcher, nil
}

// tradeEventsSliceToMap converts trade events slice to map to make comparison not order sensitive.
// This is needed because trade events slice is built from a map whose order is not guaranteeded.
func tradeEventsSliceToMap(slice []*pb.TradeEvent) map[pb.TradeEvent]struct{} {
	m := map[pb.TradeEvent]struct{}{}
	for _, e := range slice {
		m[*e] = struct{}{}
	}
	return m
}
