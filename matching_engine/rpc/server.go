package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/catortiger/exchange/matching_engine/matcher"
	mpb "github.com/catortiger/exchange/matching_engine/matcher/proto"
	spb "github.com/catortiger/exchange/matching_engine/rpc/proto"
	"github.com/catortiger/exchange/matching_engine/rpc/service"
	"google.golang.org/grpc"
)

var (
	port        = flag.String("port", ":50051", "Port.")
	tradingPair = flag.String("trading_pair", "BTC_USD", "Trading pair.")
)

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("net.Listen() = %v", err)
	}

	tp := mpb.TradingPair(mpb.TradingPair_value[*tradingPair])
	matcher := matcher.New(tp)

	srv := service.New(matcher)

	grpcServer := grpc.NewServer()
	spb.RegisterMatcherServiceServer(grpcServer, srv)
	grpcServer.Serve(lis)
}
