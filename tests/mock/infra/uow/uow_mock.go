package uow_mock

import (
	"github.com/andreis3/stores-ms/internal/infra/uow/interfaces"
	"github.com/andreis3/stores-ms/internal/util"
)

type UnitOfWorkMock struct {
	TXMock               any
	RepositoryMocks      map[string]iuow.RepositoryFactory
	RegisterFunc         func(name string, callback iuow.RepositoryFactory)
	GetRepositoryFunc    func(name string) any
	DoFunc               func(callback func(uow iuow.IUnitOfWork) *util.ValidationError) *util.ValidationError
	RollbackFunc         func() *util.ValidationError
	CommitOrRollbackFunc func() *util.ValidationError
}

func (u *UnitOfWorkMock) Register(name string, callback iuow.RepositoryFactory) {
	u.RepositoryMocks[name] = callback
	u.RegisterFunc(name, callback)
}

func (u *UnitOfWorkMock) GetRepository(name string) any {
	repo := u.RepositoryMocks[name](u.TXMock)
	u.GetRepositoryFunc(name)
	return repo
}

func (u *UnitOfWorkMock) Do(callback func(uow iuow.IUnitOfWork) *util.ValidationError) *util.ValidationError {
	callback(u)
	return u.DoFunc(callback)
}

func (u *UnitOfWorkMock) Rollback() *util.ValidationError {
	return u.RollbackFunc()
}

func (u *UnitOfWorkMock) CommitOrRollback() *util.ValidationError {
	return u.CommitOrRollbackFunc()
}
