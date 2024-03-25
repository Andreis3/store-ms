package repo_group

import (
	"github.com/andreis3/stores-ms/internal/domain/entity/group"
	"github.com/andreis3/stores-ms/internal/util"
)

type GroupModel struct {
	ID        string `db:"id"`
	GroupName string `db:"group_name"`
	Code      string `db:"code"`
	Status    string `db:"status"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}

func MapperGroupModel(group entity_group.Group) *GroupModel {
	return &GroupModel{
		ID:        group.ID,
		GroupName: group.GroupName,
		Code:      group.Code,
		Status:    group.Status.Status,
		CreatedAt: util.FormatDate(),
		UpdatedAt: util.FormatDate(),
	}
}
