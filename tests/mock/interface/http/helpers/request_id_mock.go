package helpers_mock

type RequestIDMock struct {
	GenerateFunc func() string
}

func (r *RequestIDMock) Generate() string {
	return r.GenerateFunc()
}
