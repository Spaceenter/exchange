package service

import (
	"context"
	"time"

	"github.com/catortiger/exchange/matching_engine/matcher"
	mpb "github.com/catortiger/exchange/matching_engine/matcher/proto"
	pb "github.com/catortiger/exchange/matching_engine/rpc/proto"
)

type MatcherService struct {
	matcherMap map[mpb.TradingPair]matcher.Interface
}

func New(matchers []matcher.Interface) *MatcherService {
	matcherMap := map[mpb.TradingPair]matcher.Interface{}
	for _, matcher := range matchers {
		matcherMap[matcher.TradingPair()] = matcher
	}
	return &MatcherService{matcherMap: matcherMap}
}

func (s *MatcherService) GetOrderBook(ctx context.Context,
	in *pb.GetOrderBookRequest) (*pb.GetOrderBookResponse, error) {
	out := new(pb.GetOrderBookResponse)
	matcher := s.matcherMap[in.GetTradingPair()]
	var err error
	out.OrderBook, err = matcher.OrderBook(time.Now())
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (s *MatcherService) CreateOrder(ctx context.Context,
	in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	out := new(pb.CreateOrderResponse)
	matcher := s.matcherMap[in.GetTradingPair()]
	var err error
	out.TradeEvents, out.OrderBookEvents, err = matcher.CreateOrder(in.GetOrder())
	if err != nil {
		return nil, err
	}
	return out, nil
}
