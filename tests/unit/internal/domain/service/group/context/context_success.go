package context_group_service_test

import (
	"context"
	"github.com/andreis3/stores-ms/internal/infra/repository/postgres/group"
	"github.com/andreis3/stores-ms/internal/infra/uow/interfaces"
	"github.com/andreis3/stores-ms/internal/util"
	"github.com/andreis3/stores-ms/tests/mock/infra/repository/postgres/group"
	"github.com/andreis3/stores-ms/tests/mock/infra/uow"
)

func ContextGroupServiceSuccess() *uow_mock.UnitOfWorkMock {
	var unitOfWork *uow_mock.UnitOfWorkMock
	var mapRegister []uow_mock.RegisterRepository
	var groupRepositoryMock *repo_group_mock.GroupRepositoryMock

	mapRegister = make([]uow_mock.RegisterRepository, 0)
	groupRepositoryMock = &repo_group_mock.GroupRepositoryMock{
		InsertGroupFunc: func(group repo_group.GroupModel) (string, *util.ValidationError) {
			return "1", nil
		},
	}
	mapRegister = []uow_mock.RegisterRepository{
		{
			Key:  util.GROUP_REPOSITORY_KEY,
			Repo: groupRepositoryMock,
		},
	}
	unitOfWork = &uow_mock.UnitOfWorkMock{
		DoFunc: func(ctx context.Context, callback func(iuow.IUnitOfWork) *util.ValidationError) *util.ValidationError {
			return nil
		},
		RegisterFunc: func(name string, callback iuow.RepositoryFactory) {
		},
		GetRepositoryFunc: func(name string) any {
			return nil
		},
		CommitOrRollbackFunc: func() *util.ValidationError {
			return nil
		},
		RollbackFunc: func() *util.ValidationError {
			return nil
		},
	}
	uow := uow_mock.NewProxyUnitOfWorkMock(unitOfWork, mapRegister)
	return uow
}
