package imetric

import "context"

type IPrometheusAdapter interface {
	CounterRequestHttpStatusCode(ctx context.Context, router string, statusCode int)
	HistogramRequestDuration(ctx context.Context, router string, statusCode int, duration float64)
}
