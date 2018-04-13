package matcher

import (
	"fmt"
	"testing"

	pb "github.com/spaceenter/exchange/proto"
)

func TestAddLimitOrder(t *testing.T) {
	matcher := New(pb.TradingPair_BTC_USD)
	for _, c := range []struct {
		order               *pb.Order
		wantTradeEvents     []*pb.TradeEvent
		wantOrderBookEvents []*pb.OrderBookEvent
	}{
		/*
			{
				&pb.Order{
					OrderId:   "a",
					OrderTime: testutil.protoTimeEarly,
				},
			},
		*/
	} {
		fmt.Print(matcher)
		fmt.Print(c)
	}
}
