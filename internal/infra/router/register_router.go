package router

import (
	"fmt"
	"net/http"
)

type RegisterRouter struct{}

func NewRegisterRouter() *RegisterRouter {
	return &RegisterRouter{}
}

func (r *RegisterRouter) Register(app *http.ServeMux, router []map[string]any) {
	for _, route := range router {
		fmt.Println("Registering route: ", route["method"], route["path"], route["description"])
		app.HandleFunc(
			fmt.Sprintf("%s %s", route["method"],
				route["path"]), route["handler"].(func(http.ResponseWriter, *http.Request)),
		)
	}
}
