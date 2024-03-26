package main

import (
	"net/http"

	"github.com/andreis3/stores-ms/cmd/configs"
	"github.com/andreis3/stores-ms/internal/infra/adapters/database/postgres"
	make_command "github.com/andreis3/stores-ms/internal/infra/make/command"
	"github.com/andreis3/stores-ms/internal/infra/router"
	group_controller "github.com/andreis3/stores-ms/internal/interface/http/group"
	"github.com/andreis3/stores-ms/internal/interface/http/stores"
)

func main() {
	mux := http.NewServeMux()
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	pool := postgres.NewPostgresDB(*configs)

	registerRouter := router.NewRegisterRouter()

	storesController := stores.NewStoresController()

	groupController := make_command.MakeCommandGroup(pool.DB)

	storesRouter := stores.NewStoresRouter(storesController)

	groupRouter := group_controller.NewGroupRouter(groupController)

	router.NewRouter(mux, registerRouter, storesRouter, groupRouter).ApiRoutes()

	http.ListenAndServe("0.0.0.0:8080", mux)

}
