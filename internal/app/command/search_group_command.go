package command

import (
	iservices "github.com/andreis3/stores-ms/internal/domain/services/interfaces"
	"github.com/andreis3/stores-ms/internal/interfaces/http/controllers/group/dto"
	"github.com/andreis3/stores-ms/internal/util"
)

type SearchGroupCommand struct {
	GroupService iservices.ISearchGroupService
}

func NewSearchGroupCommand(service iservices.ISearchGroupService) *SearchGroupCommand {
	return &SearchGroupCommand{
		GroupService: service,
	}
}

func (c *SearchGroupCommand) Execute(id string) (group_dto.GroupOutputDTO, *util.ValidationError) {
	output, err := c.GroupService.SearchOneGroup(id)
	if err != nil {
		return group_dto.GroupOutputDTO{}, err
	}
	return output, err
}
