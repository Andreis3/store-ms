package group_command

import "github.com/andreis3/stores-ms/internal/domain/entity/group"

type IInsertGroupCommand interface {
	Execute(data entity_group.Group) (entity_group.Group, error)
}
