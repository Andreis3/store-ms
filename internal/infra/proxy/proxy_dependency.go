package proxy

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/andreis3/stores-ms/internal/infra/adapters/database/interfaces"
	"github.com/andreis3/stores-ms/internal/infra/common/logger"
	"github.com/andreis3/stores-ms/internal/infra/make/controller"
	"github.com/andreis3/stores-ms/internal/infra/router"
	"github.com/andreis3/stores-ms/internal/interface/http/controllers/group/group_router"
	"github.com/andreis3/stores-ms/internal/interface/http/controllers/stores"
)

func ProxyDependency(mux *http.ServeMux, postgres idatabase.IDatabase, logger *logger.Logger) {
	registerRouter := router.NewRegisterRouter(logger)
	storesController := stores_controller.NewStoresController()
	groupController := make_controller.MakeCreateGroupController(postgres.InstanceDB().(*pgxpool.Pool))
	storesRouter := stores_controller.NewStoresRouter(storesController)
	groupRouter := group_router.NewGroupRouter(groupController)
	router.NewRouter(mux, registerRouter, storesRouter, groupRouter).ApiRoutes()
}
