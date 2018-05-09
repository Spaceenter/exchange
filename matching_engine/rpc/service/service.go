package service

import (
	"context"
	"time"

	"github.com/catortiger/exchange/matching_engine/matcher"
	pb "github.com/catortiger/exchange/matching_engine/rpc/proto"
)

type MatcherService struct {
	// TODO: Multiple matcher instances - one for each trading pair.
	matcher matcher.Interface
}

func New(matcher matcher.Interface) *MatcherService {
	return &MatcherService{matcher: matcher}
}

func (s *MatcherService) GetOrderBook(ctx context.Context,
	in *pb.GetOrderBookRequest) (*pb.GetOrderBookResponse, error) {
	out := new(pb.GetOrderBookResponse)
	var err error
	out.OrderBook, err = s.matcher.OrderBook(time.Now())
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (s *MatcherService) CreateOrder(ctx context.Context,
	in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	out := new(pb.CreateOrderResponse)
	var err error
	out.TradeEvents, out.OrderBookEvents, err = s.matcher.CreateOrder(in.GetOrder())
	if err != nil {
		return nil, err
	}
	return out, nil
}
