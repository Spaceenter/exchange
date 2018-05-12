package routes

import controller "github.com/CatOrTiger/exchange/web_service/rest/api_controller"

// UserRoutes ...
type UserRoutes struct {
	*Routes
}

//SetupRoute ...
func (r *UserRoutes) SetupRoute() {
	r.AddRoutes([]Route{
		Route{
			"HelloWorld",
			//You can handle more than just GET requests here, but for this tutorial we'll just do GETs
			[]string{"GET"},
			"/hello",
			// We defined HelloWorldHandler in Part1
			&ContextedHandler{&server, controller.UserControllerFactory, HelloWorldHandler},
			"user:CreateUser",
		},
		Route{
			"GoodbyeWorld",
			[]string{"GET"},
			"/goodbye",
			// GoodbyeWorldHandler is defined outside the gist :)
			&ContextedHandler{&server, controller.UserControllerFactory, GoodbyeWorldHandler},
			"user:CreateUser",
		}})
}
