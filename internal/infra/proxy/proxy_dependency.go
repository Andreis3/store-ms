package proxy

import (
	"github.com/andreis3/stores-ms/internal/infra/adapters/database/interfaces"
	"github.com/andreis3/stores-ms/internal/infra/common/logger"
	"github.com/andreis3/stores-ms/internal/infra/common/metrics/prometheus"
	"github.com/andreis3/stores-ms/internal/infra/make/controllers"
	"github.com/andreis3/stores-ms/internal/infra/routes"
	"github.com/andreis3/stores-ms/internal/interfaces/http/controllers/group/routes"
	"github.com/andreis3/stores-ms/internal/interfaces/http/controllers/stores"
	"github.com/go-chi/chi/v5"
)

func ProxyDependency(mux *chi.Mux, postgres idatabase.IDatabase, logger *logger.Logger) {
	registerRouter := routes.NewRegisterRoutes(logger)
	prometheus := metric_prometheus.NewPrometheusAdapter()
	storesController := stores_controller.NewStoresController()
	storesRouter := stores_controller.NewStoresRouter(storesController)
	createGroupController := make_controller.MakeCreateGroupController(postgres, prometheus)
	searchGroupController := make_controller.MakeSearchGroupController(postgres, prometheus)
	groupRouter := group_routes.NewGroupRoutes(createGroupController, searchGroupController)
	routes.NewRoutes(mux, registerRouter, storesRouter, groupRouter).RegisterRoutes()
}
