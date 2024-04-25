//go:build unit
// +build unit

package group_command_test

import (
	"github.com/andreis3/stores-ms/internal/interface/http/controller/group/dto"
	"github.com/andreis3/stores-ms/internal/util"
	"github.com/andreis3/stores-ms/tests/mock/domain/service/group_service"
)

func ContextInsertSuccess() *group_service_mock.InsertGroupServiceMock {
	groupServiceMock := new(group_service_mock.InsertGroupServiceMock)
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

	groupServiceMock.On(group_service_mock.InsertGroup, groupInputDTO).Return(groupOutputDTO, (*util.ValidationError)(nil))
	return groupServiceMock
}

func ContextInsertReturnErrorGroupServiceInsertGroup() *group_service_mock.InsertGroupServiceMock {

	groupServiceMock := new(group_service_mock.InsertGroupServiceMock)
	groupInputDTO := group_dto.GroupInputDTO{
		Name:   "Group 1",
		Code:   "G1",
		Status: "active",
	}
	groupOutputDTO := group_dto.GroupOutputDTO{}
	err := &util.ValidationError{
		Code:        "PIDB-235",
		Status:      500,
		ClientError: []string{"Internal Server Error"},
		LogError:    []string{"Insert group error"},
	}

	groupServiceMock.On(group_service_mock.InsertGroup, groupInputDTO).Return(groupOutputDTO, err)
	return groupServiceMock
}
