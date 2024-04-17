package repo_group

import (
	"time"

	"github.com/andreis3/stores-ms/internal/domain/entity/group"
	"github.com/andreis3/stores-ms/internal/util"
)

type GroupModel struct {
	ID        *string    `db:"id"`
	Name      *string    `db:"name"`
	Code      *string    `db:"code"`
	Status    *string    `db:"status"`
	CreatedAt *time.Time `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
}

func MapperGroupModel(group entity_group.Group) *GroupModel {
	dateTime := util.FormatDateTime()
	return &GroupModel{
		ID:        &group.ID,
		Name:      &group.Name,
		Code:      &group.Code,
		Status:    &group.Status.Status,
		CreatedAt: &dateTime,
		UpdatedAt: &dateTime,
	}
}
