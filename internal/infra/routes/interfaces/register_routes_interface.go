package iroutes

import (
	"github.com/go-chi/chi/v5"

	"github.com/andreis3/stores-ms/internal/util"
)

type IRegisterRoutes interface {
	Register(app *chi.Mux, router util.RouteType)
}
