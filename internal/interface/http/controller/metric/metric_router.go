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
