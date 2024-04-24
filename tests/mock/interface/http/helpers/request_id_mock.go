package helpers_mock

import "github.com/stretchr/testify/mock"

const (
	Generate = "Generate"
)

type RequestIDMock struct {
	mock.Mock
}

func (r *RequestIDMock) Generate() string {
	args := r.Called()
	return args.String(0)
}
