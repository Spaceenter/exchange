package main

import (
	"flag"
	"log"
	"net"
	"strings"

	"github.com/catortiger/exchange/matching_engine/matcher"
	mpb "github.com/catortiger/exchange/matching_engine/matcher/proto"
	spb "github.com/catortiger/exchange/matching_engine/rpc/proto"
	"github.com/catortiger/exchange/matching_engine/rpc/service"
	"google.golang.org/grpc"
)

var (
	port         = flag.String("port", ":50051", "Port.")
	tradingPairs = flag.String("trading_pairs", "BTC_USD,BTC_THB", "Trading pairs separated by comma.")
)

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", *port)
	if err != nil {
		log.Fatalf("net.Listen() = %v", err)
	}

	var matchers []matcher.Interface
	for _, tp := range strings.Split(*tradingPairs, ",") {
		// TODO: Filter out unknown trading pairs.
		matchers = append(matchers, matcher.New(mpb.TradingPair(mpb.TradingPair_value[tp])))
	}
	srv := service.New(matchers)

	grpcServer := grpc.NewServer()
	spb.RegisterMatcherServiceServer(grpcServer, srv)
	grpcServer.Serve(lis)
}
