package group_service

import "github.com/andreis3/stores-ms/internal/domain/entity/group"

type IInsertGroupService interface {
	InsertGroup(data group.Group) (group.Group, error)
}
