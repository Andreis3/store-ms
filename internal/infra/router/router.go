package router

import (
	"net/http"

	"github.com/andreis3/stores-ms/internal/infra/router/interfaces"
	group_controller "github.com/andreis3/stores-ms/internal/interface/http/group/interfaces"
	stores "github.com/andreis3/stores-ms/internal/interface/http/stores/interfaces"
)

type Router struct {
	router         *http.ServeMux
	registerRouter interfaces.IRegisterRouter
	storesRouter   stores.IStoresRouter
	groupRouter    group_controller.IGroupRouter
}

func NewRouter(server *http.ServeMux,
	registerRouter interfaces.IRegisterRouter,
	storesRouter stores.IStoresRouter,
	groupRouter group_controller.IGroupRouter) *Router {
	return &Router{
		router:         server,
		registerRouter: registerRouter,
		storesRouter:   storesRouter,
		groupRouter:    groupRouter,
	}
}

func (r *Router) ApiRoutes() {
	r.registerRouter.Register(r.router, r.storesRouter.StoresRoutes())
	r.registerRouter.Register(r.router, r.groupRouter.GroupRoutes())
}
