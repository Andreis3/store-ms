package uuid_mock

import "github.com/stretchr/testify/mock"

const (
	Generate = "Generate"
)

type UUIDMock struct {
	mock.Mock
}

func (r *UUIDMock) Generate() string {
	args := r.Called()
	return args.String(0)
}
