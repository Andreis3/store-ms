package metric_router

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/andreis3/stores-ms/internal/util"
)

type Router struct {
}

func NewMetricRouter() *Router {
	return &Router{}
}
func (r *Router) MetricRoutes() util.RouterType {
	return util.RouterType{
		{
			Method:      http.MethodGet,
			Path:        "/metrics",
			Controller:  promhttp.Handler(),
			Description: "Metrics Prometheus",
			Type:        util.HANDLER,
		},
	}
}
