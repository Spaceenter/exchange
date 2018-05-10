package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/catortiger/exchange/store"
	"github.com/catortiger/exchange/web_service/rest"
	"github.com/gorilla/mux"
)

var (
	port           = flag.String("port", ":8081", "Port.")
	dataSourceName = flag.String("data_source_name", "", "Data source name.")
)

func main() {
	store, err := store.New(*dataSourceName)
	if err != nil {
		log.Fatalf("store.New() = %v", err)
	}

	ws := rest.New(store)

	router := mux.NewRouter()
	router.HandleFunc("/user", ws.CreateUser).Methods("POST")

	log.Fatal(http.ListenAndServe(*port, router))
}
