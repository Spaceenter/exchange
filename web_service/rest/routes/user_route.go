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
			"get_user",
			[]string{"GET"},
			"/users/{key}",
			&ContextedHandler{&server, controller.GetUser},
		},
		Route{
			"get_users",
			[]string{"GET"},
			"/users",
			&ContextedHandler{&server, controller.GetUsers},
		},
		Route{
			"create_user",
			[]string{"POST"},
			"/users",
			&ContextedHandler{&server, controller.CreateUser},
		}})
}
