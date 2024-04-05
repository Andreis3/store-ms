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

func (r *RegisterRouter) Register(app *http.ServeMux, router []map[string]any) {
	message, info := "Registering Route ", "ROUTER"
	for _, route := range router {
		switch route[util.TYPE].(string) {
		case util.HANDLER:
			r.logger.Info(message, info, fmt.Sprintf("%s %s %s", route[util.METHOD], route[util.PATH], route[util.DESCRIPTION]))
			app.Handle(route[util.PATH].(string), route[util.CONTROLLER].(http.Handler))
		case util.HANDLER_FUNC:
			r.logger.Info(message, info, fmt.Sprintf("%s %s %s", route[util.METHOD], route[util.PATH], route[util.DESCRIPTION]))
			app.HandleFunc(
				fmt.Sprintf("%s %s", route[util.METHOD],
					route[util.PATH]), route[util.CONTROLLER].(func(http.ResponseWriter, *http.Request)))
		}
	}
}
