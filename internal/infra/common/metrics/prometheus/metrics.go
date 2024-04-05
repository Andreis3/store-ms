package prometheus

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/prometheus"
	api "go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/sdk/metric"
)

type PrometheusAdapter struct {
	counter   api.Int64Counter
	histogram api.Float64Histogram
}

func NewPrometheusAdapter() *PrometheusAdapter {
	exporter, _ := prometheus.New()
	provider := metric.NewMeterProvider(metric.WithReader(exporter))
	meter := provider.Meter("store-ms")
	counter, _ := meter.Int64Counter("proxy_requests_total", api.WithDescription("Total number of proxy requests"))
	histogram, _ := meter.Float64Histogram("request_duration_seconds", api.WithDescription("Request duration in seconds"))

	return &PrometheusAdapter{
		counter:   counter,
		histogram: histogram,
	}
}

func (p *PrometheusAdapter) CounterRequestHttpStatusCode(ctx context.Context, statusCode int) {
	opt := api.WithAttributes(
		attribute.Key("status_code").Int(statusCode),
	)
	p.counter.Add(ctx, 1, opt)
}

func (p *PrometheusAdapter) HistogramRequestDuration(ctx context.Context, groupName string, duration float64) {
	opt := api.WithAttributes(
		attribute.Key("group_name").String(groupName),
	)
	p.histogram.Record(ctx, duration, opt)
}
