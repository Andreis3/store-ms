package group_command

import (
	"github.com/andreis3/stores-ms/internal/domain/service/group_service/interfaces"
	group_dto "github.com/andreis3/stores-ms/internal/interface/http/group/dto"
	"github.com/andreis3/stores-ms/internal/util"
)

type InsertGroupCommand struct {
	GroupService group_service.IInsertGroupService
}

func NewInsertGroupCommand(service group_service.IInsertGroupService) *InsertGroupCommand {
	return &InsertGroupCommand{
		GroupService: service,
	}
}

func (c *InsertGroupCommand) Execute(data group_dto.GroupInputDTO) (group_dto.GroupOutputDTO, *util.ValidationError) {
	group, err := c.GroupService.InsertGroup(data)
	return group, err
}
