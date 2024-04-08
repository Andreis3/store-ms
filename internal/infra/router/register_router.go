package router

import (
	"fmt"
	ilogger "github.com/andreis3/stores-ms/internal/infra/common/logger/interfaces"
	"github.com/andreis3/stores-ms/internal/util"
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
func (r *RegisterRouter) Register(serverMux *http.ServeMux, router util.RouterType) {
	message, info := "[RegisterRouter] ", "MAPPED_ROUTER"
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
