package routes

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"

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
func (r *RegisterRoutes) Register(serverMux *chi.Mux, router util.RouteType) {
	message, info := "[RegisterRoutes] ", "MAPPED_ROUTER"
	for _, route := range router {
		switch route.Type {
		case util.HANDLER:
			switch len(route.Middlewares) > 0 {
			case true:
				r.logger.Info(message, info, fmt.Sprintf("%s %s - %s", route.Method, route.Path, route.Description))
				serverMux.With(route.Middlewares...).Handle(route.Path, route.Controller.(http.Handler))
			default:
				r.logger.Info(message, info, fmt.Sprintf("%s %s - %s", route.Method, route.Path, route.Description))
				serverMux.Handle(route.Path, route.Controller.(http.Handler))
			}

		case util.HANDLER_FUNC:
			switch len(route.Middlewares) > 0 {
			case true:
				r.logger.Info(message, info, fmt.Sprintf("%s %s - %s", route.Method, route.Path, route.Description))
				serverMux.With(route.Middlewares...).HandleFunc(
					route.Path, route.Controller.(func(http.ResponseWriter, *http.Request)))
			default:
				r.logger.Info(message, info, fmt.Sprintf("%s %s - %s", route.Method, route.Path, route.Description))
				serverMux.HandleFunc(
					route.Path, route.Controller.(func(http.ResponseWriter, *http.Request)))
			}
		}
	}
}
