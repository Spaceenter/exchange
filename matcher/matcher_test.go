package matcher

import (
	"reflect"
	"testing"

	pb "github.com/spaceenter/exchange/proto"
	"github.com/spaceenter/exchange/testutil"
)

func TestAddLimitOrder(t *testing.T) {
	matcher := New(pb.TradingPair_BTC_USD)
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
			t.Errorf("SubmitOrder() = %s", err)
			continue
		}
		if gotTradeEvents != nil {
			t.Errorf("SubmitOrder() TradeEvents = %s, want nil", gotTradeEvents)
		}
		if !reflect.DeepEqual(gotOrderBookEvents, c.wantOrderBookEvents) {
			t.Errorf("SubmitOrder() OrderBookEvents = %s, want %s", gotOrderBookEvents,
				c.wantOrderBookEvents)
		}
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
	gotOrderBook, err := matcher.OrderBook(testutil.TimeLate)
	if err != nil {
		t.Fatalf("OrderBook() = %s", err)
	}
	if !reflect.DeepEqual(gotOrderBook, wantOrderBook) {
		t.Errorf("OrderBook() = %s, want %s", gotOrderBook, wantOrderBook)
	}
}
