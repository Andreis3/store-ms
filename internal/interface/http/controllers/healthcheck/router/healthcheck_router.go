package healthcheck_router

import (
	"net/http"

	"github.com/andreis3/stores-ms/internal/interface/http/controllers/healthcheck/controller"
	"github.com/andreis3/stores-ms/internal/util"
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
			Controller:  healthcheck_controller.HealthCheck,
			Description: "Health Check",
			Type:        util.HANDLER_FUNC,
		},
	}
}
