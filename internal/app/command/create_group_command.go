package command

import (
	iservices "github.com/andreis3/stores-ms/internal/domain/services/interfaces"
	"github.com/andreis3/stores-ms/internal/interfaces/http/controllers/group/dto"
	"github.com/andreis3/stores-ms/internal/util"
)

type CreateGroupCommand struct {
	GroupService iservices.ICreateGroupService
}

func NewCreateGroupCommand(service iservices.ICreateGroupService) *CreateGroupCommand {
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
