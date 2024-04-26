//go:build unit
// +build unit

package group_service_test

import (
	"net/http"

	"github.com/stretchr/testify/mock"

	"github.com/andreis3/stores-ms/internal/infra/repository/postgres/group"
	"github.com/andreis3/stores-ms/internal/infra/uow/interfaces"
	"github.com/andreis3/stores-ms/internal/util"
	"github.com/andreis3/stores-ms/tests/mock/infra/common/uuid_mock"
	"github.com/andreis3/stores-ms/tests/mock/infra/repository/postgres/group"
	"github.com/andreis3/stores-ms/tests/mock/infra/uow"
)

func ContextInsertSuccess(groupRepositoryMock *repo_group_mock.GroupRepositoryMock, uuidMock *uuid_mock.UUIDMock) *uow_mock.UnitOfWorkMock {
	unitOfWork := new(uow_mock.UnitOfWorkMock)

	groupOutputDTO := &repo_group.GroupModel{
		ID:        util.StringToPointer("123"),
		Name:      util.StringToPointer("Group 1"),
		Code:      util.StringToPointer("G1"),
		Status:    util.StringToPointer("active"),
		CreatedAt: util.FormatDateStringToPointerTime("2021-01-01T00:00:00Z"),
		UpdatedAt: util.FormatDateStringToPointerTime("2021-01-01T00:00:00Z"),
	}

	groupRepositoryMock.On(repo_group_mock.SelectOneGroupByNameAndCode, mock.Anything, mock.Anything).Return(&repo_group.GroupModel{}, (*util.ValidationError)(nil))
	groupRepositoryMock.On(repo_group_mock.InsertGroup, mock.Anything).Return(groupOutputDTO, &util.ValidationError{})
	uuidMock.On(uuid_mock.Generate).Return("123")

	unitOfWork.On(uow_mock.Do, mock.AnythingOfType(uow_mock.DoParamFunc)).Return((*util.ValidationError)(nil)).Run(func(args mock.Arguments) {
		callback := args.Get(0).(func(iuow.IUnitOfWork) *util.ValidationError)
		callback(unitOfWork)
	}).Once()

	unitOfWork.On(uow_mock.GetRepository, util.GROUP_REPOSITORY_KEY).Return(groupRepositoryMock)

	return unitOfWork
}
func ContextInsertReturnErrorGroupRepositoryInsertGroup(groupRepositoryMock *repo_group_mock.GroupRepositoryMock, uuidMock *uuid_mock.UUIDMock) *uow_mock.UnitOfWorkMock {
	unitOfWork := new(uow_mock.UnitOfWorkMock)
	err := &util.ValidationError{
		Code:        "PIDB-235",
		Status:      http.StatusInternalServerError,
		LogError:    []string{"Insert group error"},
		ClientError: []string{"Internal Server Error"},
	}
	groupRepositoryMock.On(repo_group_mock.SelectOneGroupByNameAndCode, mock.Anything, mock.Anything).Return(&repo_group.GroupModel{}, &util.ValidationError{})
	groupRepositoryMock.On(repo_group_mock.InsertGroup, mock.Anything).Return("", err)
	uuidMock.On(uuid_mock.Generate).Return("123")

	unitOfWork.On(uow_mock.Do, mock.AnythingOfType(uow_mock.DoParamFunc)).Return(err).Run(func(args mock.Arguments) {
		callback := args.Get(0).(func(iuow.IUnitOfWork) *util.ValidationError)
		callback(unitOfWork)
	}).Once()

	unitOfWork.On(uow_mock.GetRepository, util.GROUP_REPOSITORY_KEY).Return(groupRepositoryMock)

	return unitOfWork
}
func ContextInsertReturnErrorWhenCommitCommandUow(groupRepositoryMock *repo_group_mock.GroupRepositoryMock, uuidMock *uuid_mock.UUIDMock) *uow_mock.UnitOfWorkMock {
	unitOfWork := new(uow_mock.UnitOfWorkMock)
	err := &util.ValidationError{
		Code:        "PIDB-235",
		Status:      http.StatusInternalServerError,
		LogError:    []string{"Commit error"},
		ClientError: []string{"Internal Server Error"},
	}

	groupRepositoryMock.On(repo_group_mock.SelectOneGroupByNameAndCode, mock.Anything, mock.Anything).Return(&repo_group.GroupModel{}, &util.ValidationError{})
	groupRepositoryMock.On(repo_group_mock.InsertGroup, mock.Anything).Return("1", nil)
	uuidMock.On(uuid_mock.Generate).Return("123")

	unitOfWork.On(uow_mock.Do, mock.AnythingOfType(uow_mock.DoParamFunc)).Return(err).Run(func(args mock.Arguments) {
		callback := args.Get(0).(func(iuow.IUnitOfWork) *util.ValidationError)
		callback(unitOfWork)
	}).Once()

	unitOfWork.On(uow_mock.CommitOrRollback).Return(err).Run(func(args mock.Arguments) {
		callback := args.Get(0).(func() *util.ValidationError)
		callback()
	}).Once()

	unitOfWork.On(uow_mock.GetRepository, util.GROUP_REPOSITORY_KEY).Return(groupRepositoryMock)

	return unitOfWork
}
func ContextInsertReturnErrorWhenRollbackCommandUow(groupRepositoryMock *repo_group_mock.GroupRepositoryMock, uuidMock *uuid_mock.UUIDMock) *uow_mock.UnitOfWorkMock {
	unitOfWork := new(uow_mock.UnitOfWorkMock)

	err := &util.ValidationError{
		Code:        "PIDB-235",
		Status:      http.StatusInternalServerError,
		LogError:    []string{"Rollback error"},
		ClientError: []string{"Internal Server Error"},
	}

	groupRepositoryMock.On(repo_group_mock.SelectOneGroupByNameAndCode, mock.Anything, mock.Anything).Return(&repo_group.GroupModel{}, nil)
	groupRepositoryMock.On(repo_group_mock.InsertGroup, mock.Anything).Return("1", nil)
	uuidMock.On(uuid_mock.Generate).Return("123")

	unitOfWork.On(uow_mock.Do, mock.AnythingOfType(uow_mock.DoParamFunc)).Return(err).Run(func(args mock.Arguments) {
		callback := args.Get(0).(func(iuow.IUnitOfWork) *util.ValidationError)
		callback(unitOfWork)
	}).Once()

	unitOfWork.On(uow_mock.CommitOrRollback).Return(nil).Run(func(args mock.Arguments) {
		callback := args.Get(0).(func() *util.ValidationError)
		callback()
	}).Once()

	unitOfWork.On(uow_mock.Rollback).Return(err).Run(func(args mock.Arguments) {
		callback := args.Get(0).(func() *util.ValidationError)
		callback()
	}).Once()

	unitOfWork.On(uow_mock.GetRepository, util.GROUP_REPOSITORY_KEY).Return(&groupRepositoryMock)

	return unitOfWork
}
func ContextInsertReturnErrorWhenSelectOneGroupByNameAndCode(groupRepositoryMock *repo_group_mock.GroupRepositoryMock, uuidMock *uuid_mock.UUIDMock) *uow_mock.UnitOfWorkMock {
	unitOfWork := new(uow_mock.UnitOfWorkMock)

	err := &util.ValidationError{
		Code:        "PIDB-235",
		Status:      http.StatusInternalServerError,
		LogError:    []string{"Select group error"},
		ClientError: []string{"Internal Server Error"},
	}

	groupRepositoryMock.On(repo_group_mock.SelectOneGroupByNameAndCode, mock.Anything, mock.Anything).Return(&repo_group.GroupModel{}, err)
	groupRepositoryMock.On(repo_group_mock.InsertGroup, mock.Anything).Return("", &util.ValidationError{})
	uuidMock.On(uuid_mock.Generate).Return("123")

	unitOfWork.On(uow_mock.Do, mock.AnythingOfType(uow_mock.DoParamFunc)).Return(err).Run(func(args mock.Arguments) {
		callback := args.Get(0).(func(iuow.IUnitOfWork) *util.ValidationError)
		callback(unitOfWork)
	}).Once()

	unitOfWork.On(uow_mock.GetRepository, util.GROUP_REPOSITORY_KEY).Return(groupRepositoryMock)

	return unitOfWork

}
func ContextInsertReturnErrorWhenSelectOneGroupByNameAndCodeReturnGroup(groupRepositoryMock *repo_group_mock.GroupRepositoryMock, uuidMock *uuid_mock.UUIDMock) *uow_mock.UnitOfWorkMock {
	unitOfWork := new(uow_mock.UnitOfWorkMock)

	err := &util.ValidationError{
		Code:        "VBR-0002",
		LogError:    []string{"Group already exists"},
		ClientError: []string{"Group already exists"},
		Status:      http.StatusBadRequest,
	}

	model := &repo_group.GroupModel{
		ID:     util.StringToPointer("1"),
		Name:   util.StringToPointer("Group 1"),
		Code:   util.StringToPointer("G1"),
		Status: util.StringToPointer("active"),
	}

	groupRepositoryMock.On(repo_group_mock.SelectOneGroupByNameAndCode, mock.Anything, mock.Anything).Return(model, &util.ValidationError{})
	groupRepositoryMock.On(repo_group_mock.InsertGroup, mock.Anything).Return("", &util.ValidationError{})
	uuidMock.On(uuid_mock.Generate).Return("123")
	unitOfWork.On(uow_mock.Do, mock.AnythingOfType(uow_mock.DoParamFunc)).Return(err).Run(func(args mock.Arguments) {
		callback := args.Get(0).(func(iuow.IUnitOfWork) *util.ValidationError)
		callback(unitOfWork)

	}).Once()

	unitOfWork.On(uow_mock.GetRepository, util.GROUP_REPOSITORY_KEY).Return(groupRepositoryMock)

	return unitOfWork

}
