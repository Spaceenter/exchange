package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/CatOrTiger/exchange/store"
	"github.com/CatOrTiger/exchange/web_service/rest"
	"github.com/CatOrTiger/exchange/web_service/rest/routes"
)

var (
	port           = flag.String("port", ":8081", "Port.")
	dataSourceName = flag.String("data_source_name", "", "Data source name.") //TODO fix later
)

// initDb() ininal database the store should get connection and connection poll
func initDb() *store.Store {
	store, err := store.New(*dataSourceName)
	if err != nil {
		// log.Fatalf("store.New() = %v", err)
		return nil
	}
	return store
}

//TODO do your redis pool here add a new function

func main() {
	//inital global verabials add init
	server := rest.WebService{DB: initDb()}

	//configeaution
	//test connections
	//test mouduls
	mx := routes.InitRouter(server)

	log.Fatal(http.ListenAndServe(*port, mx))
}
