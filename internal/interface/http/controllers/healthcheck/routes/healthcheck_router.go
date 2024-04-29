package healthcheck_routes

import (
	"net/http"

	"github.com/andreis3/stores-ms/internal/interface/http/controllers/healthcheck/controller"
	"github.com/andreis3/stores-ms/internal/util"
)

type HealthCheckRouter struct{}

func NewHealthCheckRoutes() *HealthCheckRouter {
	return &HealthCheckRouter{}
}
func (r *HealthCheckRouter) HealthCheckRoutes() util.RouteType {
	return util.RouteType{
		{
			Method:      http.MethodGet,
			Path:        "/healthcheck",
			Controller:  healthcheck_controller.HealthCheck,
			Description: "Health Check",
			Type:        util.HANDLER_FUNC,
			Middlewares: []func(http.Handler) http.Handler{},
		},
	}
}
