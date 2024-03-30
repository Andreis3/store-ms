package main

import (
	"net/http"

	"github.com/andreis3/stores-ms/internal/infra/common/logger"
	"github.com/andreis3/stores-ms/internal/interface/http/stores"

	"github.com/andreis3/stores-ms/cmd/configs"
	"github.com/andreis3/stores-ms/internal/infra/adapters/database/postgres"
	"github.com/andreis3/stores-ms/internal/infra/make/controller"
	"github.com/andreis3/stores-ms/internal/infra/router"
	"github.com/andreis3/stores-ms/internal/interface/http/group"
)

func main() {
	mux := http.NewServeMux()
	conf, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	pool := postgres.NewPostgresDB(*conf)
	logger := logger.NewLogger()

	registerRouter := router.NewRegisterRouter(logger)

	storesController := stores_controller.NewStoresController()

	groupController := make_controller.MakeControllerGroup(pool.DB)

	storesRouter := stores_controller.NewStoresRouter(storesController)

	groupRouter := group_controller.NewGroupRouter(groupController)

	router.NewRouter(mux, registerRouter, storesRouter, groupRouter).ApiRoutes()

	err = http.ListenAndServe(":"+conf.ServerPort, mux)
	if err != nil {
		panic(err)
	}

}
