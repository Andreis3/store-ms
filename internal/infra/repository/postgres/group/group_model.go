package group

import "github.com/andreis3/stores-ms/internal/domain/entity/group"

type GroupModel struct {
	ID        string `db:"id"`
	GroupName string `db:"group_name"`
	Code      string `db:"code"`
	Status    string `db:"status"`
}

func MapperGroupModel(group entity_group.Group) *GroupModel {
	return &GroupModel{
		ID:        group.ID,
		GroupName: group.GroupName,
		Code:      group.Code,
		Status:    group.Status.Status,
	}
}
