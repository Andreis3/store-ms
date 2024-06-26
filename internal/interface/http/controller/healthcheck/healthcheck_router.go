package healthcheck_controller

import (
	"github.com/andreis3/stores-ms/internal/util"
	"net/http"
)

type Router struct{}

func NewHealthCheckRouter() *Router {
	return &Router{}
}
func (r *Router) HealthCheckRoutes() util.RouterType {
	return util.RouterType{
		{
			Method:      http.MethodGet,
			Path:        "/healthcheck",
			Controller:  HealthCheck,
			Description: "Health Check",
			Type:        util.HANDLER_FUNC,
		},
	}
}
