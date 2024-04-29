package router

import (
	"net/http"

	"github.com/andreis3/stores-ms/internal/interface/http/controllers/healthcheck/router"
	"github.com/andreis3/stores-ms/internal/interface/http/controllers/metric"

	"github.com/andreis3/stores-ms/internal/infra/router/interfaces"
	"github.com/andreis3/stores-ms/internal/interface/http/controllers/group/interfaces"
	"github.com/andreis3/stores-ms/internal/interface/http/controllers/stores/interfaces"
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
	r.registerRouter.Register(r.serverMux, healthcheck_router.NewHealthCheckRouter().HealthCheckRoutes())
	r.registerRouter.Register(r.serverMux, metric_controller.NewMetricRouter().MetricRoutes())
}
