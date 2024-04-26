package repo_group_mock

import (
	"github.com/stretchr/testify/mock"

	entity_group "github.com/andreis3/stores-ms/internal/domain/entity/group"
	"github.com/andreis3/stores-ms/internal/infra/repository/postgres/group"
	"github.com/andreis3/stores-ms/internal/util"
)

const (
	InsertGroup                 = "InsertGroup"
	SelectOneGroupByNameAndCode = "SelectOneGroupByNameAndCode"
)

type GroupRepositoryMock struct {
	mock.Mock
}

func (g *GroupRepositoryMock) InsertGroup(data entity_group.Group) (*repo_group.GroupModel, *util.ValidationError) {
	args := g.Called(data)
	return args.Get(0).(*repo_group.GroupModel), args.Get(1).(*util.ValidationError)
}
func (g *GroupRepositoryMock) SelectOneGroupByNameAndCode(name, code string) (*repo_group.GroupModel, *util.ValidationError) {
	args := g.Called(name, code)
	return args.Get(0).(*repo_group.GroupModel), args.Get(1).(*util.ValidationError)
}
