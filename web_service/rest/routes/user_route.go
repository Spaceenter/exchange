package routes

import controller "github.com/CatOrTiger/exchange/web_service/rest/controller"

// UserRoutes ...
type UserRoutes struct {
	*Routes
}

//SetupRoute ...
func (r *UserRoutes) SetupRoute() {
	r.AddRoutes([]Route{
		Route{
			"get user",
			//You can handle more than just GET requests here, but for this tutorial we'll just do GETs
			[]string{"GET"},
			"/users/{key}",
			// We defined HelloWorldHandler in Part1
			&ContextedHandler{&server, controller.GetUser},
		},
		Route{
			"get users",
			[]string{"GET"},
			"/users",
			// GoodbyeWorldHandler is defined outside the gist :)
			&ContextedHandler{&server, controller.GetUsers},
		},
		Route{
			"create user",
			[]string{"POST"},
			"/users",
			// GoodbyeWorldHandler is defined outside the gist :)
			&ContextedHandler{&server, controller.CreateUser},
		}})
}
