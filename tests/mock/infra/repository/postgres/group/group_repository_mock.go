package repo_group_mock

import (
	"github.com/andreis3/stores-ms/internal/infra/repository/postgres/group"
	"github.com/andreis3/stores-ms/internal/util"
)

type GroupRepositoryMock struct {
	InsertGroupFunc                 func(group repo_group.GroupModel) (string, *util.ValidationError)
	SelectOneGroupByNameAndCodeFunc func(name, code string) (*repo_group.GroupModel, *util.ValidationError)
}

func NewGroupRepositoryMock() *GroupRepositoryMock {
	return &GroupRepositoryMock{}
}

func (g *GroupRepositoryMock) InsertGroup(group repo_group.GroupModel) (string, *util.ValidationError) {
	return g.InsertGroupFunc(group)
}
func (g *GroupRepositoryMock) SelectOneGroupByNameAndCode(name, code string) (*repo_group.GroupModel, *util.ValidationError) {
	return g.SelectOneGroupByNameAndCodeFunc(name, code)
}
