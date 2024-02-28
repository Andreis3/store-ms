package main

import (
	"net/http"

	"github.com/andreis3/stores-ms/internal/infra/router"
	"github.com/andreis3/stores-ms/internal/interface/http/stores"
)

func main() {
	mux := http.NewServeMux()

	registerRouter := router.NewRegisterRouter()

	storesController := stores.NewStoresController()

	storesRouter := stores.NewStoresRouter(storesController)

	router.NewRouter(mux, registerRouter, storesRouter).ApiRoutes()

	http.ListenAndServe("0.0.0.0:8080", mux)

}
