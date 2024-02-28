package router

import (
	"net/http"

	"github.com/andreis3/stores-ms/internal/infra/router/interfaces"
	stores "github.com/andreis3/stores-ms/internal/interface/http/stores/interfaces"
)

type Router struct {
	router         *http.ServeMux
	registerRouter interfaces.IRegisterRouter
	storesRouter   stores.IStoresRouter
}

func NewRouter(server *http.ServeMux,
	registerRouter interfaces.IRegisterRouter,
	storesRouter stores.IStoresRouter) *Router {
	return &Router{
		router:         server,
		registerRouter: registerRouter,
		storesRouter:   storesRouter,
	}
}

func (r *Router) ApiRoutes() {
	r.registerRouter.Register(r.router, r.storesRouter.StoresRoutes())
}
