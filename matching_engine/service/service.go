package service

import (
	"context"

	pb "github.com/spaceenter/exchange/matching_engine/service/proto"
)

type MatcherService struct{}

func New() *MatcherService {
	return nil
}

func (s *MatcherService) GetOrderBook(context.Context,
	*pb.GetOrderBookRequest) (*pb.GetOrderBookResponse, error) {
	return nil, nil
}
