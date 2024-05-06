package metric_prometheus_mock

import (
	"context"

	"github.com/stretchr/testify/mock"
)

const (
	CounterRequestHttpStatusCode = "CounterRequestHttpStatusCode"
	HistogramRequestDuration     = "HistogramRequestDuration"
)

type PrometheusAdapterMock struct {
	mock.Mock
}

func (p *PrometheusAdapterMock) CounterRequestHttpStatusCode(ctx context.Context, router string, statusCode int) {
	p.Called(ctx, router, statusCode)
}

func (p *PrometheusAdapterMock) HistogramRequestDuration(ctx context.Context, router string, statusCode int, duration float64) {
	p.Called(ctx, router, statusCode, duration)
}

func (p *PrometheusAdapterMock) HistogramInstructionTableDuration(ctx context.Context, database, table, method string, duration float64) {
	p.Called(ctx, database, table, method, duration)
}
