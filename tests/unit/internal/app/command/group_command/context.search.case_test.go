//go:build unit
// +build unit

package group_command_test

import (
	"github.com/andreis3/stores-ms/internal/interfaces/http/controllers/group/dto"
	"github.com/andreis3/stores-ms/internal/util"
	"github.com/andreis3/stores-ms/tests/mock/domain/service/group_service"
)

func ContextSearchSuccess() *group_service_mock.SearchGroupServiceMock {
	searchServiceMock := new(group_service_mock.SearchGroupServiceMock)
	id := "1"

	groupOutputDTO := group_dto.GroupOutputDTO{
		ID:        "1",
		Name:      "Group 1",
		Code:      "G1",
		Status:    "active",
		CreatedAt: "2021-01-01T00:00:00Z",
		UpdatedAt: "2021-01-01T00:00:00Z",
	}

	searchServiceMock.On(group_service_mock.SearchOneGroup, id).Return(groupOutputDTO, (*util.ValidationError)(nil))
	return searchServiceMock
}

func ContextSearchReturnErrorGroupServiceInsertGroup() *group_service_mock.SearchGroupServiceMock {

	searchServiceMock := new(group_service_mock.SearchGroupServiceMock)
	id := "1"

	groupOutputDTO := group_dto.GroupOutputDTO{}
	err := &util.ValidationError{
		Code:        "PIDB-235",
		Status:      500,
		ClientError: []string{"Internal Server Error"},
		LogError:    []string{"Insert group error"},
	}

	searchServiceMock.On(group_service_mock.SearchOneGroup, id).Return(groupOutputDTO, err)
	return searchServiceMock
}
