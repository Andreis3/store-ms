package group_service_mock

import (
	"github.com/andreis3/stores-ms/internal/interface/http/controller/group/dto"
	"github.com/andreis3/stores-ms/internal/util"
)

type InsertGroupServiceMock struct {
	InsertGroupFunc func(data group_dto.GroupInputDTO) (group_dto.GroupOutputDTO, *util.ValidationError)
}

func (m *InsertGroupServiceMock) InsertGroup(data group_dto.GroupInputDTO) (group_dto.GroupOutputDTO, *util.ValidationError) {
	return m.InsertGroupFunc(data)
}
