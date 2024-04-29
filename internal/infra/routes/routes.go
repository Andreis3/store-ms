package routes

import (
	"net/http"

	"github.com/andreis3/stores-ms/internal/infra/routes/interfaces"
	"github.com/andreis3/stores-ms/internal/interface/http/controllers/group/interfaces"
	"github.com/andreis3/stores-ms/internal/interface/http/controllers/healthcheck/router"
	"github.com/andreis3/stores-ms/internal/interface/http/controllers/metric/router"
	"github.com/andreis3/stores-ms/internal/interface/http/controllers/stores/interfaces"
)

type Routes struct {
	serverMux      *http.ServeMux
	registerRouter iroutes.IRegisterRoutes
	storesRouter   istores_controller.IStoresRouter
	groupRouter    igroup_controller.IGroupRouter
}

func NewRoutes(serverMux *http.ServeMux,
	registerRouter iroutes.IRegisterRoutes,
	storesRouter istores_controller.IStoresRouter,
	groupRouter igroup_controller.IGroupRouter) *Routes {
	return &Routes{
		serverMux:      serverMux,
		registerRouter: registerRouter,
		storesRouter:   storesRouter,
		groupRouter:    groupRouter,
	}
}
func (r *Routes) RegisterRoutes() {
	r.registerRouter.Register(r.serverMux, r.storesRouter.StoresRoutes())
	r.registerRouter.Register(r.serverMux, r.groupRouter.GroupRoutes())
	r.registerRouter.Register(r.serverMux, healthcheck_router.NewHealthCheckRouter().HealthCheckRoutes())
	r.registerRouter.Register(r.serverMux, metric_router.NewMetricRouter().MetricRoutes())
}
