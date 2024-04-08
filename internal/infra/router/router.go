package router

import (
	"github.com/andreis3/stores-ms/internal/interface/http/controller/healthcheck"
	"github.com/andreis3/stores-ms/internal/interface/http/controller/metric"
	"net/http"

	"github.com/andreis3/stores-ms/internal/infra/router/interfaces"
	"github.com/andreis3/stores-ms/internal/interface/http/controller/group/interfaces"
	"github.com/andreis3/stores-ms/internal/interface/http/controller/stores/interfaces"
)

type Router struct {
	serverMux      *http.ServeMux
	registerRouter irouter.IRegisterRouter
	storesRouter   istores_controller.IStoresRouter
	groupRouter    igroup_controller.IGroupRouter
}

func NewRouter(serverMux *http.ServeMux,
	registerRouter irouter.IRegisterRouter,
	storesRouter istores_controller.IStoresRouter,
	groupRouter igroup_controller.IGroupRouter) *Router {
	return &Router{
		serverMux:      serverMux,
		registerRouter: registerRouter,
		storesRouter:   storesRouter,
		groupRouter:    groupRouter,
	}
}
func (r *Router) ApiRoutes() {
	r.registerRouter.Register(r.serverMux, r.storesRouter.StoresRoutes())
	r.registerRouter.Register(r.serverMux, r.groupRouter.GroupRoutes())
	r.registerRouter.Register(r.serverMux, healthcheck_controller.NewHealthCheckRouter().HealthCheckRoutes())
	r.registerRouter.Register(r.serverMux, metric_controller.NewMetricRouter().MetricRoutes())
}
