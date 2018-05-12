package routes

import (
	"io"
	"log"
	"net/http"

	"github.com/CatOrTiger/exchange/web_service/rest"
	"github.com/CatOrTiger/exchange/web_service/rest/controller"

	"github.com/CatOrTiger/exchange/web_service/rest/middleware"
	"github.com/gorilla/mux"
)

// Route ..
type Route struct {
	Name             string
	Method           []string
	Pattern          string
	ContextedHandler *ContextedHandler
	Controller       string
}

//Routes just stores our Route declarations
type Routes struct {
	routes []Route
}

// main webservice instance
var server rest.WebService

// route manager
var routeManager *Routes

type routeManagerInterface interface {
	SetupRoute()
}

// AddRoutes to route manager
func (m *Routes) AddRoutes(newRoutes []Route) {
	m.routes = append(m.routes, newRoutes...)
}

//InitRouter returns a new Gorrila Mux router
func InitRouter(c rest.WebService) *mux.Router {
	muxRouter := mux.NewRouter().StrictSlash(true)
	server = c
	routeManager = &Routes{}

	routes := []routeManagerInterface{
		&UserRoutes{routeManager},
	}

	for _, route := range routes {
		route.SetupRoute()
	}

	for _, route := range routeManager.routes {
		muxRouter.
			Methods(route.Method...).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.ContextedHandler)
	}

	muxRouter.Use(middleware.LoggingMiddleware)
	//Check all routes to make sure the users are properly authenticated
	muxRouter.Use(middleware.Authenticate)
	// muxRouter.Use(middleware.SetContentTypeText)
	return muxRouter
}

//ContextedHandler is a wrapper to provide AppContext to our Handlers
type ContextedHandler struct {
	server      *rest.WebService
	ProcessFunc func(*rest.WebService, *controller.RequestInfo, controller.Query, io.ReadCloser) (*controller.Response, error)
}

// ServeHTTP ...
func (handler *ContextedHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//TODO
	//here should be checked authentication ok.
	//parse the input header
	//parse the input body
	//parse the queries
	//need a factory to create correct controller.
	//need a binding object pass it to api controller
	info := &controller.RequestInfo{
		Protocal:   "json", // change to enum
		APIVersion: "v1",   // change to enum
	}

	status, err := handler.ProcessFunc(handler.server, info, mux.Vars(r), r.Body)
	if err != nil {
		log.Printf("HTTP %d: %q", status, err)
		switch status {
		// TODO you can handle any error that a router might return here.
		}
	}
}
