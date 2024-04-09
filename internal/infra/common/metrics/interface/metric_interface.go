package imetric

import "context"

type IMetricAdapter interface {
	CounterRequestHttpStatusCode(ctx context.Context, router string, statusCode int)
	HistogramRequestDuration(ctx context.Context, router string, statusCode int, duration float64)
}
