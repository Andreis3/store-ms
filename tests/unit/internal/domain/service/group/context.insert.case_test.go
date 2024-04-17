//go:build unit
// +build unit

package group_service_test

import (
	"context"
	"github.com/andreis3/stores-ms/internal/infra/repository/postgres/group"
	"github.com/andreis3/stores-ms/internal/infra/uow/interfaces"
	"github.com/andreis3/stores-ms/internal/util"
	"github.com/andreis3/stores-ms/tests/mock/infra/repository/postgres/group"
	"github.com/andreis3/stores-ms/tests/mock/infra/uow"
	"net/http"
)

func ContextInsertSuccess() *uow_mock.UnitOfWorkMock {
	var unitOfWork *uow_mock.UnitOfWorkMock
	var mapRegister []uow_mock.RegisterRepository
	var groupRepositoryMock *repo_group_mock.GroupRepositoryMock

	mapRegister = make([]uow_mock.RegisterRepository, 0)
	groupRepositoryMock = &repo_group_mock.GroupRepositoryMock{
		InsertGroupFunc: func(group repo_group.GroupModel) (string, *util.ValidationError) {
			return "1", nil
		},
		SelectOneGroupByNameAndCodeFunc: func(name, code string) (*repo_group.GroupModel, *util.ValidationError) {
			return &repo_group.GroupModel{}, nil
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
func ContextInsertReturnErrorGroupRepositoryInsertGroup() *uow_mock.UnitOfWorkMock {
	var unitOfWork *uow_mock.UnitOfWorkMock
	var mapRegister []uow_mock.RegisterRepository
	var groupRepositoryMock *repo_group_mock.GroupRepositoryMock

	mapRegister = make([]uow_mock.RegisterRepository, 0)
	groupRepositoryMock = &repo_group_mock.GroupRepositoryMock{
		InsertGroupFunc: func(group repo_group.GroupModel) (string, *util.ValidationError) {
			return "", &util.ValidationError{
				Code:        "PIDB-235",
				Status:      http.StatusInternalServerError,
				LogError:    []string{"Insert group error"},
				ClientError: []string{"Internal Server Error"},
			}
		},
		SelectOneGroupByNameAndCodeFunc: func(name, code string) (*repo_group.GroupModel, *util.ValidationError) {
			return &repo_group.GroupModel{}, nil
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
			return &util.ValidationError{
				Code:        "PIDB-235",
				Status:      http.StatusInternalServerError,
				LogError:    []string{"Insert group error"},
				ClientError: []string{"Internal Server Error"},
			}
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
func ContextInsertReturnErrorWhenCommitCommandUow() *uow_mock.UnitOfWorkMock {
	var unitOfWork *uow_mock.UnitOfWorkMock
	var mapRegister []uow_mock.RegisterRepository
	var groupRepositoryMock *repo_group_mock.GroupRepositoryMock

	mapRegister = make([]uow_mock.RegisterRepository, 0)
	groupRepositoryMock = &repo_group_mock.GroupRepositoryMock{
		InsertGroupFunc: func(group repo_group.GroupModel) (string, *util.ValidationError) {
			return "1", nil
		},
		SelectOneGroupByNameAndCodeFunc: func(name, code string) (*repo_group.GroupModel, *util.ValidationError) {
			return &repo_group.GroupModel{}, nil
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
			return &util.ValidationError{
				Code:        "PIDB-235",
				Status:      http.StatusInternalServerError,
				LogError:    []string{"Commit error"},
				ClientError: []string{"Internal Server Error"},
			}
		},
		RegisterFunc: func(name string, callback iuow.RepositoryFactory) {
		},
		GetRepositoryFunc: func(name string) any {
			return nil
		},
		CommitOrRollbackFunc: func() *util.ValidationError {
			return &util.ValidationError{
				Code:        "PIDB-235",
				Status:      http.StatusInternalServerError,
				LogError:    []string{"Commit error"},
				ClientError: []string{"Internal Server Error"},
			}
		},
		RollbackFunc: func() *util.ValidationError {
			return nil
		},
	}
	uow := uow_mock.NewProxyUnitOfWorkMock(unitOfWork, mapRegister)
	return uow
}
func ContextInsertReturnErrorWhenRollbackCommandUow() *uow_mock.UnitOfWorkMock {
	var unitOfWork *uow_mock.UnitOfWorkMock
	var mapRegister []uow_mock.RegisterRepository
	var groupRepositoryMock *repo_group_mock.GroupRepositoryMock

	mapRegister = make([]uow_mock.RegisterRepository, 0)
	groupRepositoryMock = &repo_group_mock.GroupRepositoryMock{
		InsertGroupFunc: func(group repo_group.GroupModel) (string, *util.ValidationError) {
			return "1", nil
		},
		SelectOneGroupByNameAndCodeFunc: func(name, code string) (*repo_group.GroupModel, *util.ValidationError) {
			return &repo_group.GroupModel{}, nil
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
			return &util.ValidationError{
				Code:        "PIDB-235",
				Status:      http.StatusInternalServerError,
				LogError:    []string{"Rollback error"},
				ClientError: []string{"Internal Server Error"},
			}
		},
		RegisterFunc: func(name string, callback iuow.RepositoryFactory) {
		},
		GetRepositoryFunc: func(name string) any {
			return nil
		},
		CommitOrRollbackFunc: func() *util.ValidationError {
			return &util.ValidationError{
				Code:        "PIDB-235",
				Status:      http.StatusInternalServerError,
				LogError:    []string{"Rollback error"},
				ClientError: []string{"Internal Server Error"},
			}
		},
		RollbackFunc: func() *util.ValidationError {
			return &util.ValidationError{
				Code:        "PIDB-235",
				Status:      http.StatusInternalServerError,
				LogError:    []string{"Rollback error"},
				ClientError: []string{"Internal Server Error"},
			}
		},
	}
	uow := uow_mock.NewProxyUnitOfWorkMock(unitOfWork, mapRegister)
	return uow
}
func ContextInsertReturnErrorWhenSelectOneGroupByNameAndCode() *uow_mock.UnitOfWorkMock {
	var unitOfWork *uow_mock.UnitOfWorkMock
	var mapRegister []uow_mock.RegisterRepository
	var groupRepositoryMock *repo_group_mock.GroupRepositoryMock

	mapRegister = make([]uow_mock.RegisterRepository, 0)
	groupRepositoryMock = &repo_group_mock.GroupRepositoryMock{
		InsertGroupFunc: func(group repo_group.GroupModel) (string, *util.ValidationError) {
			return "1", nil
		},
		SelectOneGroupByNameAndCodeFunc: func(name, code string) (*repo_group.GroupModel, *util.ValidationError) {
			return nil, &util.ValidationError{
				Code:        "PIDB-235",
				Status:      http.StatusInternalServerError,
				LogError:    []string{"Select group error"},
				ClientError: []string{"Internal Server Error"},
			}
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
			return &util.ValidationError{
				Code:        "PIDB-235",
				Status:      http.StatusInternalServerError,
				LogError:    []string{"Select group error"},
				ClientError: []string{"Internal Server Error"},
			}
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
func ContextInsertReturnErrorWhenSelectOneGroupByNameAndCodeReturnGroup() *uow_mock.UnitOfWorkMock {
	var unitOfWork *uow_mock.UnitOfWorkMock
	var mapRegister []uow_mock.RegisterRepository
	var groupRepositoryMock *repo_group_mock.GroupRepositoryMock

	mapRegister = make([]uow_mock.RegisterRepository, 0)
	groupRepositoryMock = &repo_group_mock.GroupRepositoryMock{
		InsertGroupFunc: func(group repo_group.GroupModel) (string, *util.ValidationError) {
			return "1", nil
		},
		SelectOneGroupByNameAndCodeFunc: func(name, code string) (*repo_group.GroupModel, *util.ValidationError) {
			date := util.FormatDateTime()
			return &repo_group.GroupModel{
				ID:        util.StringToPointer("1"),
				Name:      util.StringToPointer("Group 1"),
				Code:      util.StringToPointer("G1"),
				CreatedAt: &date,
				UpdatedAt: &date,
			}, nil
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
			return &util.ValidationError{
				Code:        "VBR-0002",
				LogError:    []string{"Group already exists"},
				ClientError: []string{"Group already exists"},
				Status:      http.StatusBadRequest,
			}
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
