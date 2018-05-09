package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/spaceenter/exchange/matching_engine/matcher"
	mpb "github.com/spaceenter/exchange/matching_engine/matcher/proto"
	"github.com/spaceenter/exchange/matching_engine/service"
	spb "github.com/spaceenter/exchange/matching_engine/service/proto"
	"github.com/spaceenter/exchange/store"
	"google.golang.org/grpc"
)

var (
	dataSourceName = flag.String("data_source_name", "", "Data source name.")
	port           = flag.String("port", ":50051", "Port.")
	tradingPair    = flag.String("trading_pair", "BTC_USD", "Trading pair.")
)

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("net.Listen() = %v", err)
	}

	tp := mpb.TradingPair(mpb.TradingPair_value[*tradingPair])
	matcher := matcher.New(tp)

	store, err := store.New(*dataSourceName)
	if err != nil {
		log.Fatalf("store.New() = %v", err)
	}

	srv := service.New(matcher, store)

	grpcServer := grpc.NewServer()
	spb.RegisterMatcherServiceServer(grpcServer, srv)
	grpcServer.Serve(lis)
}
