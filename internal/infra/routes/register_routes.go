package routes

import (
	"fmt"
	"net/http"

	"github.com/andreis3/stores-ms/internal/infra/common/logger/interfaces"
	"github.com/andreis3/stores-ms/internal/util"
)

type RegisterRoutes struct {
	logger ilogger.ILogger
}

func NewRegisterRoutes(logger ilogger.ILogger) *RegisterRoutes {
	return &RegisterRoutes{
		logger: logger,
	}
}
func (r *RegisterRoutes) Register(serverMux *http.ServeMux, router util.RouterType) {
	message, info := "[RegisterRoutes] ", "MAPPED_ROUTER"
	for _, route := range router {
		switch route.Type {
		case util.HANDLER:
			r.logger.Info(message, info, fmt.Sprintf("%s %s - %s", route.Method, route.Path, route.Description))
			serverMux.Handle(route.Path, route.Controller.(http.Handler))
		case util.HANDLER_FUNC:
			r.logger.Info(message, info, fmt.Sprintf("%s %s - %s", route.Method, route.Path, route.Description))
			serverMux.HandleFunc(
				fmt.Sprintf("%s %s", route.Method,
					route.Path), route.Controller.(func(http.ResponseWriter, *http.Request)))
		}
	}
}
