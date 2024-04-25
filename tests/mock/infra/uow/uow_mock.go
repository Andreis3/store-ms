package uow_mock

import (
	"github.com/stretchr/testify/mock"

	"github.com/andreis3/stores-ms/internal/infra/uow/interfaces"
	"github.com/andreis3/stores-ms/internal/util"
)

type UnitOfWorkMock struct {
	mock.Mock
}

func (u *UnitOfWorkMock) Register(name string, callback iuow.RepositoryFactory) {
	u.Called(name, callback)
}

func (u *UnitOfWorkMock) GetRepository(name string) any {
	args := u.Called(name)
	return args.Get(0)
}

func (u *UnitOfWorkMock) Do(callback func(uow iuow.IUnitOfWork) *util.ValidationError) *util.ValidationError {
	args := u.Called(callback)
	return args.Get(0).(*util.ValidationError)
}

func (u *UnitOfWorkMock) Rollback() *util.ValidationError {
	args := u.Called()
	return args.Get(0).(*util.ValidationError)
}

func (u *UnitOfWorkMock) CommitOrRollback() *util.ValidationError {
	args := u.Called()
	return args.Get(0).(*util.ValidationError)
}
