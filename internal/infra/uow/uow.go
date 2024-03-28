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
	repo := u.Repositories[name](u.TX)
	return repo
}
func (u *UnitOfWork) Do(ctx context.Context, callback func(uow iuow.IUnitOfWork) *util.ValidationError) *util.ValidationError {
	if u.TX != nil {
		return &util.ValidationError{
			LogError:    []string{"transaction already exists"},
			ClientError: []string{"Internal Server Error"},
			Status:      http.StatusInternalServerError}
	}
	tx, err := u.DB.Begin(ctx)
	if err != nil {
		return &util.ValidationError{
			LogError:    []string{err.Error()},
			ClientError: []string{"Internal Server Error"},
			Status:      http.StatusInternalServerError}
	}
	u.TX = tx
	errCB := callback(u)
	if errCB.ExistError() {
		errRb := u.Rollback()
		if errRb != nil {
			return &util.ValidationError{
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
			LogError:    []string{"transaction not exists"},
			ClientError: []string{"Internal Server Error"},
			Status:      http.StatusInternalServerError,
		}
	}
	err := u.TX.Rollback(context.Background())
	if err != nil {
		return &util.ValidationError{
			LogError:    []string{err.Error()},
			ClientError: []string{"Internal Server Error"},
			Status:      http.StatusInternalServerError,
		}
	}
	u.TX = nil
	return nil
}
func (u *UnitOfWork) CommitOrRollback() *util.ValidationError {
	if u.TX == nil {
		return nil
	}
	if err := u.TX.Commit(context.Background()); err != nil {
		if errRB := u.Rollback(); errRB != nil {
			return errRB
		}
		return &util.ValidationError{
			LogError:    []string{err.Error()},
			ClientError: []string{"Internal Server Error"},
			Status:      http.StatusInternalServerError}
	}
	u.TX = nil
	return nil
}
