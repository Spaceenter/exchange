package service

import (
	"context"
	"time"

	"github.com/spaceenter/exchange/matching_engine/matcher"
	pb "github.com/spaceenter/exchange/matching_engine/service/proto"
)

type MatcherService struct {
	matcher matcher.MatcherInterface
}

func New(matcher matcher.MatcherInterface) *MatcherService {
	return &MatcherService{
		matcher: matcher,
	}
}

func (s *MatcherService) GetOrderBook(ctx context.Context,
	in *pb.GetOrderBookRequest) (*pb.GetOrderBookResponse, error) {
	out := new(pb.GetOrderBookResponse)
	var err error
	out.OrderBook, err = s.matcher.OrderBook(time.Now())
	return out, err
}

func (s *MatcherService) CreateOrder(ctx context.Context,
	in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	out := new(pb.CreateOrderResponse)
	var err error
	out.TradeEvents, out.OrderBookEvents, err = s.matcher.CreateOrder(in.GetOrder())
	return out, err
}
