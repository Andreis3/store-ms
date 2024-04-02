package router

import (
	"net/http"

	"github.com/andreis3/stores-ms/internal/infra/router/interfaces"
	"github.com/andreis3/stores-ms/internal/interface/http/controller/group/interfaces"
	stores "github.com/andreis3/stores-ms/internal/interface/http/controller/stores/interfaces"
)

type Router struct {
	router         *http.ServeMux
	registerRouter irouter.IRegisterRouter
	storesRouter   stores.IStoresRouter
	groupRouter    igroup_controller.IGroupRouter
}

func NewRouter(server *http.ServeMux,
	registerRouter irouter.IRegisterRouter,
	storesRouter stores.IStoresRouter,
	groupRouter igroup_controller.IGroupRouter) *Router {
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
