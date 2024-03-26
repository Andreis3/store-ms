package irepo_group

import "github.com/andreis3/stores-ms/internal/infra/repository/postgres/group"

type IGroupRepository interface {
	InsertGroup(group repo_group.GroupModel) (string, error)
}
