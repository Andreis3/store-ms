package group_command_mock

import (
	"github.com/andreis3/stores-ms/internal/interface/http/controller/group/dto"
	"github.com/andreis3/stores-ms/internal/util"
)

type InsertGroupCommandMock struct {
	ExecuteFunc func(data group_dto.GroupInputDTO) (group_dto.GroupOutputDTO, *util.ValidationError)
}

func (i *InsertGroupCommandMock) Execute(data group_dto.GroupInputDTO) (group_dto.GroupOutputDTO, *util.ValidationError) {
	return i.ExecuteFunc(data)
}
