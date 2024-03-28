package main

import (
	"github.com/andreis3/stores-ms/internal/infra/common/logger"
	"net/http"

	"github.com/andreis3/stores-ms/cmd/configs"
	"github.com/andreis3/stores-ms/internal/infra/adapters/database/postgres"
	make_controller "github.com/andreis3/stores-ms/internal/infra/make/controller"
	"github.com/andreis3/stores-ms/internal/infra/router"
	group_controller "github.com/andreis3/stores-ms/internal/interface/http/group"
	"github.com/andreis3/stores-ms/internal/interface/http/stores"
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

	storesController := stores.NewStoresController()

	groupController := make_controller.MakeControllerGroup(pool.DB)

	storesRouter := stores.NewStoresRouter(storesController)

	groupRouter := group_controller.NewGroupRouter(groupController)

	router.NewRouter(mux, registerRouter, storesRouter, groupRouter).ApiRoutes()

	err = http.ListenAndServe(":"+conf.ServerPort, mux)
	if err != nil {
		panic(err)
	}

}
