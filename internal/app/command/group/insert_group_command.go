package group_command

import (
	"github.com/andreis3/stores-ms/internal/domain/service/group/interfaces"
	"github.com/andreis3/stores-ms/internal/interface/http/controllers/group/dto"
	"github.com/andreis3/stores-ms/internal/util"
)

type InsertGroupCommand struct {
	GroupService igroup_service.IInsertGroupService
}

func NewInsertGroupCommand(service igroup_service.IInsertGroupService) *InsertGroupCommand {
	return &InsertGroupCommand{
		GroupService: service,
	}
}

func (c *InsertGroupCommand) Execute(data group_dto.GroupInputDTO) (group_dto.GroupOutputDTO, *util.ValidationError) {
	group := data.MapperInputDtoToEntity()
	output, err := c.GroupService.InsertGroup(*group)
	if err != nil {
		return group_dto.GroupOutputDTO{}, err
	}
	return output, err
}
