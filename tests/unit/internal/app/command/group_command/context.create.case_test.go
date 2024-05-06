//go:build unit
// +build unit

package group_command_test

import (
	"github.com/andreis3/stores-ms/internal/interfaces/http/controllers/group/dto"
	"github.com/andreis3/stores-ms/internal/util"
	"github.com/andreis3/stores-ms/tests/mock/domain/service/group_service"
)

func ContextCreateSuccess() *group_service_mock.CreateGroupServiceMock {
	groupServiceMock := new(group_service_mock.CreateGroupServiceMock)
	groupInputDTO := group_dto.GroupInputDTO{
		Name:   "Group 1",
		Code:   "G1",
		Status: "active",
	}
	groupOutputDTO := group_dto.GroupOutputDTO{
		ID:        "1",
		Name:      "Group 1",
		Code:      "G1",
		Status:    "active",
		CreatedAt: "2021-01-01T00:00:00Z",
		UpdatedAt: "2021-01-01T00:00:00Z",
	}

	group := groupInputDTO.MapperInputDtoToEntity()

	groupServiceMock.On(group_service_mock.CreateGroup, *group).Return(groupOutputDTO, (*util.ValidationError)(nil))
	return groupServiceMock
}

func ContextCreateReturnErrorGroupServiceCreateGroup() *group_service_mock.CreateGroupServiceMock {

	groupServiceMock := new(group_service_mock.CreateGroupServiceMock)
	groupInputDTO := group_dto.GroupInputDTO{
		Name:   "Group 1",
		Code:   "G1",
		Status: "active",
	}

	gourp := groupInputDTO.MapperInputDtoToEntity()
	groupOutputDTO := group_dto.GroupOutputDTO{}
	err := &util.ValidationError{
		Code:        "PIDB-235",
		Status:      500,
		ClientError: []string{"Internal Server Error"},
		LogError:    []string{"Insert group error"},
	}

	groupServiceMock.On(group_service_mock.CreateGroup, *gourp).Return(groupOutputDTO, err)
	return groupServiceMock
}
