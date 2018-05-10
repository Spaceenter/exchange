package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/catortiger/exchange/web_service/rest"
	"github.com/gorilla/mux"
)

var (
	port = flag.String("port", ":8081", "Port.")
)

func main() {
	ws := rest.New()
	router := mux.NewRouter()
	router.HandleFunc("/user", ws.CreateUser).Methods("POST")
	log.Fatal(http.ListenAndServe(*port, router))
}
