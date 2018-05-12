package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/CatOrTiger/exchange/web_service/rest"
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
	FactoryFunc func(apiVersion string, protocal string) interface{}
	//ContextedHandlerFunc is the interface which our Handlers will implement
	ContextedHandlerFunc func(*rest.WebService, http.ResponseWriter, *http.Request) (int, error)
}

type RequestObject struct {
	Name  string
	Value int
}

type testInterface interface {
}

func (handler *ContextedHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// here should be checked authentication ok.

	//parse the input header
	//parse the input body
	//parse the queries

	//get request body depens on `get post put`.
	// test := controller.UsersController
	// type := reflect.TypeOf(testInterface)

	vars := mux.Vars(r)

	for _, v := range vars {
		println(v)
	}

	println(r.Body)

	for _, h := range r.Header {
		println(h)
	}

	// m := controller.UsersController{}
	// meth := reflect.ValueOf(m).MethodByName("CreateUser")
	// meth.Call(nil)

	//need a factory to create correct controller.

	//need a binding object pass it to api controller
	status, err := handler.ContextedHandlerFunc(handler.server, w, r)
	if err != nil {
		log.Printf("HTTP %d: %q", status, err)
		switch status {
		// TODO you can handle any error that a router might return here.
		// I use SendGrid to send an email anytime a 503 occurs :)z
		}
	}
}

//HelloWorldHandler just prints Hello World on a GET request
func HelloWorldHandler(s *rest.WebService, res http.ResponseWriter, req *http.Request) (int, error) {
	//So in this handler we now have the context provided
	fmt.Fprint(res, "Hello World")
	return http.StatusOK, nil
}

// GoodbyeWorldHandler ...
func GoodbyeWorldHandler(s *rest.WebService, res http.ResponseWriter, req *http.Request) (int, error) {
	//So in this handler we now have the context provided
	fmt.Fprint(res, "Hello World 1")
	return http.StatusOK, nil
}
