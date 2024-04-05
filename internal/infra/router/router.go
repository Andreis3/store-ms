package router

import (
	metric_controller "github.com/andreis3/stores-ms/internal/interface/http/controller/metric"
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
	metric         metric_controller.Router
}

func NewRouter(server *http.ServeMux,
	registerRouter irouter.IRegisterRouter,
	storesRouter stores.IStoresRouter,
	groupRouter igroup_controller.IGroupRouter) *Router {
	metric := metric_controller.NewMetricRouter()
	return &Router{
		router:         server,
		registerRouter: registerRouter,
		storesRouter:   storesRouter,
		groupRouter:    groupRouter,
		metric:         *metric,
	}
}

func (r *Router) ApiRoutes() {
	r.registerRouter.Register(r.router, r.storesRouter.StoresRoutes())
	r.registerRouter.Register(r.router, r.groupRouter.GroupRoutes())
	r.registerRouter.Register(r.router, r.metric.MetricRoutes())
}
