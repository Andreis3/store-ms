package router

import (
	"fmt"
	ilogger "github.com/andreis3/stores-ms/internal/infra/common/logger/interfaces"
	"net/http"
)

type RegisterRouter struct {
	logger ilogger.ILogger
}

func NewRegisterRouter(logger ilogger.ILogger) *RegisterRouter {
	return &RegisterRouter{
		logger: logger,
	}
}

func (r *RegisterRouter) Register(app *http.ServeMux, router []map[string]any) {
	for _, route := range router {
		r.logger.Info("Registering Route ", "ROUTER", fmt.Sprintf("%s %s %s", route["method"], route["path"], route["description"]))
		app.HandleFunc(
			fmt.Sprintf("%s %s", route["method"],
				route["path"]), route["handler"].(func(http.ResponseWriter, *http.Request)),
		)
	}
}
