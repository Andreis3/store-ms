package group_service

import "github.com/andreis3/stores-ms/internal/domain/entity/group"

type IInsertGroupService interface {
	InsertGroup(data entity_group.Group) (entity_group.Group, error)
}
