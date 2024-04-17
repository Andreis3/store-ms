package irepo_group

import (
	"github.com/andreis3/stores-ms/internal/infra/repository/postgres/group"
	"github.com/andreis3/stores-ms/internal/util"
)

type IGroupRepository interface {
	InsertGroup(group repo_group.GroupModel) (string, *util.ValidationError)
	SelectOneGroupByNameAndCode(name, code string) (*repo_group.GroupModel, *util.ValidationError)
}
