package group_command_mock

import (
	"github.com/andreis3/stores-ms/internal/interface/http/controller/group/dto"
	"github.com/andreis3/stores-ms/internal/util"
)

type InsertGroupCommandMock struct {
	FuncParamsInput  []any
	FuncParamsOutput []any
	ExecuteFunc      func(data group_dto.GroupInputDTO) (group_dto.GroupOutputDTO, *util.ValidationError)
}

func (i *InsertGroupCommandMock) Execute(data group_dto.GroupInputDTO) (group_dto.GroupOutputDTO, *util.ValidationError) {
	i.FuncParamsInput = []any{data}
	returnOne, returnTwo := i.ExecuteFunc(data)
	i.FuncParamsOutput = []any{returnOne, returnTwo}
	return returnOne, returnTwo
}
