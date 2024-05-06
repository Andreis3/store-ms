package group_command

import (
	"github.com/andreis3/stores-ms/internal/domain/service/group/interfaces"
	iuow "github.com/andreis3/stores-ms/internal/infra/uow/interfaces"
	"github.com/andreis3/stores-ms/internal/interfaces/http/controllers/group/dto"
	"github.com/andreis3/stores-ms/internal/util"
)

type SearchGroupCommand struct {
	GroupService igroup_service.ISearchGroupService
}

func NewSearchGroupCommand(service igroup_service.ISearchGroupService) *SearchGroupCommand {
	return &SearchGroupCommand{
		GroupService: service,
	}
}

func (c *SearchGroupCommand) Execute(id string, uow iuow.IUnitOfWork) (group_dto.GroupOutputDTO, *util.ValidationError) {
	output, err := c.GroupService.SearchOneGroup(id, uow)
	if err != nil {
		return group_dto.GroupOutputDTO{}, err
	}
	return output, err
}
