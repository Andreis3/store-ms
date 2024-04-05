package metric_controller

import (
	"github.com/andreis3/stores-ms/internal/util"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

type Router struct {
}

func NewMetricRouter() *Router {
	return &Router{}
}

func (r *Router) MetricRoutes() []map[string]any {
	return []map[string]any{
		{
			util.METHOD:      http.MethodGet,
			util.PATH:        "/metrics",
			util.CONTROLLER:  promhttp.Handler(),
			util.DESCRIPTION: "Metrics Prometheus",
			util.TYPE:        util.HANDLER,
		},
	}
}
