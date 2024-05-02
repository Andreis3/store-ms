package proxy

import (
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/andreis3/stores-ms/internal/infra/adapters/database/interfaces"
	"github.com/andreis3/stores-ms/internal/infra/common/logger"
	"github.com/andreis3/stores-ms/internal/infra/make/controllers"
	"github.com/andreis3/stores-ms/internal/infra/routes"
	"github.com/andreis3/stores-ms/internal/interfaces/http/controllers/group/routes"
	"github.com/andreis3/stores-ms/internal/interfaces/http/controllers/stores"
)

func ProxyDependency(mux *chi.Mux, postgres idatabase.IDatabase, logger *logger.Logger) {
	registerRouter := routes.NewRegisterRoutes(logger)
	storesController := stores_controller.NewStoresController()
	groupController := make_controller.MakeCreateGroupController(postgres.InstanceDB().(*pgxpool.Pool))
	storesRouter := stores_controller.NewStoresRouter(storesController)
	groupRouter := group_routes.NewGroupRoutes(groupController)
	routes.NewRoutes(mux, registerRouter, storesRouter, groupRouter).RegisterRoutes()
}
