package group_command

import (
	"github.com/andreis3/stores-ms/internal/domain/service/group/interfaces"
	"github.com/andreis3/stores-ms/internal/interfaces/http/controllers/group/dto"
	"github.com/andreis3/stores-ms/internal/util"
)

type CreateGroupCommand struct {
	GroupService igroup_service.ICreateGroupService
}

func NewCreateGroupCommand(service igroup_service.ICreateGroupService) *CreateGroupCommand {
	return &CreateGroupCommand{
		GroupService: service,
	}
}

func (c *CreateGroupCommand) Execute(data group_dto.GroupInputDTO) (group_dto.GroupOutputDTO, *util.ValidationError) {
	group := data.MapperInputDtoToEntity()
	output, err := c.GroupService.CreateGroup(*group)
	if err != nil {
		return group_dto.GroupOutputDTO{}, err
	}
	return output, err
}
