package iroutes

import (
	"net/http"

	"github.com/andreis3/stores-ms/internal/util"
)

type IRegisterRoutes interface {
	Register(app *http.ServeMux, router util.RouterType)
}
