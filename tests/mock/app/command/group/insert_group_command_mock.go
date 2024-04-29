package group_command_mock

import (
	"github.com/stretchr/testify/mock"

	"github.com/andreis3/stores-ms/internal/interface/http/controllers/group/dto"
	"github.com/andreis3/stores-ms/internal/util"
)

const (
	Execute = "Execute"
)

type InsertGroupCommandMock struct {
	mock.Mock
}

func (igc *InsertGroupCommandMock) Execute(data group_dto.GroupInputDTO) (group_dto.GroupOutputDTO, *util.ValidationError) {
	args := igc.Called(data)
	return args.Get(0).(group_dto.GroupOutputDTO), args.Get(1).(*util.ValidationError)
}
