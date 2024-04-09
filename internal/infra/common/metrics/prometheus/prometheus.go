package metric_prometheus

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/prometheus"
	api "go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/sdk/metric"
)

const (
	METER_NAME    = "store-ms"
	METER_VERSION = "1.0.0"
)

type PrometheusAdapter struct {
	counter   api.Int64Counter
	histogram api.Float64Histogram
}

func NewPrometheusAdapter() *PrometheusAdapter {
	exporter, _ := prometheus.New()
	provider := metric.NewMeterProvider(metric.WithReader(exporter))
	meter := provider.Meter(METER_NAME, api.WithInstrumentationVersion(METER_VERSION))
	counter, _ := meter.Int64Counter("proxy_requests_total",
		api.WithDescription("Total number of proxy requests"))
	histogram, _ := meter.Float64Histogram("request_duration_seconds",
		api.WithDescription("Request duration in seconds"),
		api.WithExplicitBucketBoundaries(5, 10, 15, 20, 30, 50, 100, 200, 300, 500, 1000, 2000, 5000, 10000, 20000, 30000))
	return &PrometheusAdapter{
		counter:   counter,
		histogram: histogram,
	}
}
func (p *PrometheusAdapter) CounterRequestHttpStatusCode(ctx context.Context, router string, statusCode int) {
	opt := api.WithAttributes(
		attribute.Key("router").String(router),
		attribute.Key("status_code").Int(statusCode),
	)
	p.counter.Add(ctx, 1, opt)
}
func (p *PrometheusAdapter) HistogramRequestDuration(ctx context.Context, router string, statusCode int, duration float64) {
	opt := api.WithAttributes(
		attribute.Key("router").String(router),
		attribute.Key("status_code").Int(statusCode),
	)
	p.histogram.Record(ctx, duration, opt)
}
