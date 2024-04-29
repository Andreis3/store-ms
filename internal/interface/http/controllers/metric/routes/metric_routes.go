package metric_router

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/andreis3/stores-ms/internal/util"
)

type MetricsRoutes struct {
}

func NewMetricRouter() *MetricsRoutes {
	return &MetricsRoutes{}
}
func (r *MetricsRoutes) MetricRoutes() util.RouteType {
	return util.RouteType{
		{
			Method:      http.MethodGet,
			Path:        "/metrics",
			Controller:  promhttp.Handler(),
			Description: "Metrics Prometheus",
			Type:        util.HANDLER,
			Middlewares: []func(http.Handler) http.Handler{},
		},
	}
}
