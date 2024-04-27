package uow

import (
	"context"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/andreis3/stores-ms/internal/infra/uow/interfaces"
	"github.com/andreis3/stores-ms/internal/util"
)

type UnitOfWork struct {
	DB           *pgxpool.Pool
	TX           pgx.Tx
	Repositories map[string]iuow.RepositoryFactory
}

func NewUnitOfWork(db *pgxpool.Pool) *UnitOfWork {
	return &UnitOfWork{
		DB:           db,
		Repositories: make(map[string]iuow.RepositoryFactory),
	}
}
func (u *UnitOfWork) Register(name string, callback iuow.RepositoryFactory) {
	u.Repositories[name] = callback
}
func (u *UnitOfWork) GetRepository(name string) any {
	ctx := context.Background()
	if u.TX == nil {
		tx, err := u.DB.Begin(ctx)
		if err != nil {
			return nil
		}
		u.TX = tx
	}
	repo := u.Repositories[name](u.TX)
	return repo
}
func (u *UnitOfWork) Do(callback func(uow iuow.IUnitOfWork) *util.ValidationError) *util.ValidationError {
	ctx := context.Background()
	if u.TX != nil {
		return &util.ValidationError{
			Code:        "PDB-0001",
			Origin:      "UnitOfWork.Do",
			LogError:    []string{"transaction already exists"},
			ClientError: []string{"Internal Server Error"},
			Status:      http.StatusInternalServerError}
	}
	tx, err := u.DB.Begin(ctx)
	if err != nil {
		return &util.ValidationError{
			Code:        "PDB-0000",
			Origin:      "UnitOfWork.Do",
			LogError:    []string{err.Error()},
			ClientError: []string{"Internal Server Error"},
			Status:      http.StatusInternalServerError}
	}
	u.TX = tx
	errCB := callback(u)
	if errCB != nil {
		errRb := u.Rollback()
		if errRb != nil {
			return &util.ValidationError{
				Code:        errRb.Code,
				Origin:      errRb.Origin,
				LogError:    append(errCB.LogError, errRb.LogError...),
				ClientError: []string{"Internal Server Error"},
				Status:      http.StatusInternalServerError}
		}
		return errCB
	}
	return u.CommitOrRollback()
}
func (u *UnitOfWork) Rollback() *util.ValidationError {
	if u.TX == nil {
		return &util.ValidationError{
			Code:        "PDB-0003",
			Origin:      "UnitOfWork.Rollback",
			LogError:    []string{"transaction not exists"},
			ClientError: []string{"Internal Server Error"},
			Status:      http.StatusInternalServerError,
		}
	}
	ctx := context.Background()
	err := u.TX.Rollback(ctx)
	if err != nil {
		return &util.ValidationError{
			Code:        "PDB-0002",
			Origin:      "UnitOfWork.Rollback",
			LogError:    []string{err.Error()},
			ClientError: []string{"Internal Server Error"},
			Status:      http.StatusInternalServerError,
		}
	}
	u.TX = nil
	return nil
}
func (u *UnitOfWork) CommitOrRollback() *util.ValidationError {
	ctx := context.Background()
	if u.TX == nil {
		return nil
	}
	if err := u.TX.Commit(ctx); err != nil {
		if errRB := u.Rollback(); errRB != nil {
			return errRB
		}
		return &util.ValidationError{
			Code:        "PDB-0004",
			Origin:      "UnitOfWork.CommitOrRollback",
			LogError:    []string{err.Error()},
			ClientError: []string{"Internal Server Error"},
			Status:      http.StatusInternalServerError}
	}
	u.TX = nil
	return nil
}
