//go:build unit
// +build unit

package group_command_test

import (
	"github.com/andreis3/stores-ms/internal/interface/http/controller/group/dto"
	"github.com/andreis3/stores-ms/internal/util"
	"github.com/andreis3/stores-ms/tests/mock/domain/service/group_service"
)

func ContextInsertSuccess() *group_service_mock.InsertGroupServiceMock {
	return &group_service_mock.InsertGroupServiceMock{
		InsertGroupFunc: func(data group_dto.GroupInputDTO) (group_dto.GroupOutputDTO, *util.ValidationError) {
			return group_dto.GroupOutputDTO{
				ID:        "1",
				Name:      "Group 1",
				Code:      "G1",
				Status:    "active",
				CreatedAt: "2021-01-01T00:00:00Z",
				UpdatedAt: "2021-01-01T00:00:00Z",
			}, nil
		},
	}
}

func ContextInsertReturnErrorGroupServiceInsertGroup() *group_service_mock.InsertGroupServiceMock {
	return &group_service_mock.InsertGroupServiceMock{
		InsertGroupFunc: func(data group_dto.GroupInputDTO) (group_dto.GroupOutputDTO, *util.ValidationError) {
			return group_dto.GroupOutputDTO{}, &util.ValidationError{
				Code:        "PIDB-235",
				Status:      500,
				ClientError: []string{"Internal Server Error"},
				LogError:    []string{"Insert group error"},
			}
		},
	}
}
