package service

import (
	"context"
	"time"

	"github.com/spaceenter/exchange/matching_engine/matcher"
	pb "github.com/spaceenter/exchange/matching_engine/service/proto"
	"github.com/spaceenter/exchange/store"
)

type MatcherService struct {
	matcher matcher.Interface
	store   store.Interface
}

func New(matcher matcher.Interface, store store.Interface) *MatcherService {
	return &MatcherService{
		matcher: matcher,
		store:   store,
	}
}

func (s *MatcherService) GetOrderBook(ctx context.Context,
	in *pb.GetOrderBookRequest) (*pb.GetOrderBookResponse, error) {
	out := new(pb.GetOrderBookResponse)
	var err error
	out.OrderBook, err = s.matcher.OrderBook(time.Now())
	if err != nil {
		return nil, err
	}
	// TODO: Add records to DB.
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
	// TODO: Add records to DB.
	return out, nil
}
