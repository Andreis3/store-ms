package group_service_mock

import (
	"github.com/stretchr/testify/mock"

	"github.com/andreis3/stores-ms/internal/interface/http/controller/group/dto"
	"github.com/andreis3/stores-ms/internal/util"
)

const (
	InsertGroup = "InsertGroup"
)

type InsertGroupServiceMock struct {
	mock.Mock
}

func (m *InsertGroupServiceMock) InsertGroup(data group_dto.GroupInputDTO) (group_dto.GroupOutputDTO, *util.ValidationError) {
	args := m.Called(data)
	return args.Get(0).(group_dto.GroupOutputDTO), args.Get(1).(*util.ValidationError)
}
