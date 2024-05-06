package group_service_mock

import (
	"github.com/stretchr/testify/mock"

	entity_group "github.com/andreis3/stores-ms/internal/domain/entity/group"
	"github.com/andreis3/stores-ms/internal/interfaces/http/controllers/group/dto"
	"github.com/andreis3/stores-ms/internal/util"
)

const (
	CreateGroup = "CreateGroup"
)

type CreateGroupServiceMock struct {
	mock.Mock
}

func (m *CreateGroupServiceMock) CreateGroup(data entity_group.Group) (group_dto.GroupOutputDTO, *util.ValidationError) {
	args := m.Called(data)
	return args.Get(0).(group_dto.GroupOutputDTO), args.Get(1).(*util.ValidationError)
}
