package irepo_group

import (
	entity_group "github.com/andreis3/stores-ms/internal/domain/entity/group"
	"github.com/andreis3/stores-ms/internal/infra/repository/postgres/group"
	"github.com/andreis3/stores-ms/internal/util"
)

type IGroupRepository interface {
	InsertGroup(data entity_group.Group) (*repo_group.GroupModel, *util.ValidationError)
	SelectOneGroupByNameAndCode(name, code string) (*repo_group.GroupModel, *util.ValidationError)
	SelectOneGroupByID(id string) (*repo_group.GroupModel, *util.ValidationError)
}
