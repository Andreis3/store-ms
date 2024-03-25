package group_command

import (
	"github.com/andreis3/stores-ms/internal/domain/entity/group"
	"github.com/andreis3/stores-ms/internal/domain/service/group_service/interfaces"
)

type InsertGroupCommand struct {
	GroupService group_service.IInsertGroupService
}

func (c InsertGroupCommand) Execute(data entity_group.Group) (entity_group.Group, error) {
	return c.GroupService.InsertGroup(data)
}
