package metric_prometheus_mock

import "context"

type PrometheusAdapterMock struct {
	CounterRequestHttpStatusCodeFunc func(ctx context.Context, router string, statusCode int)
	HistogramRequestDurationFunc     func(ctx context.Context, router string, statusCode int, duration float64)
}

func NewPrometheusAdapterMock() *PrometheusAdapterMock {
	return &PrometheusAdapterMock{}
}

func (p *PrometheusAdapterMock) CounterRequestHttpStatusCode(ctx context.Context, router string, statusCode int) {
	p.CounterRequestHttpStatusCodeFunc(ctx, router, statusCode)
}

func (p *PrometheusAdapterMock) HistogramRequestDuration(ctx context.Context, router string, statusCode int, duration float64) {
	p.HistogramRequestDurationFunc(ctx, router, statusCode, duration)
}
