package proxy

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/andreis3/stores-ms/internal/infra/adapters/database/postgres"
	"github.com/andreis3/stores-ms/internal/infra/common/configs"
	"github.com/andreis3/stores-ms/internal/infra/common/logger"
	"github.com/andreis3/stores-ms/internal/infra/make/controller"
	"github.com/andreis3/stores-ms/internal/infra/router"
	"github.com/andreis3/stores-ms/internal/interface/http/group"
	"github.com/andreis3/stores-ms/internal/interface/http/stores"
)

func ProxyDependency(mux *http.ServeMux, pool *postgres.Postgres, logger *logger.Logger, conf *configs.Conf) {
	registerRouter := router.NewRegisterRouter(logger)
	storesController := stores_controller.NewStoresController()
	groupController := make_controller.MakeControllerGroup(pool.InstanceDB().(*pgxpool.Pool))
	storesRouter := stores_controller.NewStoresRouter(storesController)
	groupRouter := group_controller.NewGroupRouter(groupController)
	router.NewRouter(mux, registerRouter, storesRouter, groupRouter).ApiRoutes()
}
