package irouter

import (
	"github.com/andreis3/stores-ms/internal/util"
	"net/http"
)

type IRegisterRouter interface {
	Register(app *http.ServeMux, router util.RouterType)
}
