package group_command_mock

import (
	"github.com/stretchr/testify/mock"

	group_dto "github.com/andreis3/stores-ms/internal/interfaces/http/controllers/group/dto"
	"github.com/andreis3/stores-ms/internal/util"
)

type SelectGroupCommandMock struct {
	mock.Mock
}

func (m *SelectGroupCommandMock) Execute(id string) (group_dto.GroupOutputDTO, *util.ValidationError) {
	ret := m.Called(id)

	var r0 group_dto.GroupOutputDTO
	if rf, ok := ret.Get(0).(func(string) group_dto.GroupOutputDTO); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(group_dto.GroupOutputDTO)
	}

	var r1 *util.ValidationError
	if rf, ok := ret.Get(1).(func(string) *util.ValidationError); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Get(1).(*util.ValidationError)
	}

	return r0, r1
}
