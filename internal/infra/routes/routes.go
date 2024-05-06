package routes

import (
	"github.com/go-chi/chi/v5"

	"github.com/andreis3/stores-ms/internal/infra/routes/interfaces"
	"github.com/andreis3/stores-ms/internal/interfaces/http/controllers/group/interfaces"
	"github.com/andreis3/stores-ms/internal/interfaces/http/controllers/healthcheck/routes"
	"github.com/andreis3/stores-ms/internal/interfaces/http/controllers/metric/routes"
	"github.com/andreis3/stores-ms/internal/interfaces/http/controllers/stores/interfaces"
)

type Routes struct {
	serverMux      *chi.Mux
	registerRouter iroutes.IRegisterRoutes
	storesRouter   istores_controller.IStoresRouter
	groupRouter    igroup_controller.IGroupRouter
}

func NewRoutes(serverMux *chi.Mux,
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
	r.registerRouter.Register(r.serverMux, healthcheck_routes.NewHealthCheckRoutes().HealthCheckRoutes())
	r.registerRouter.Register(r.serverMux, metric_router.NewMetricRouter().MetricRoutes())
}
