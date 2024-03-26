package uow

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/andreis3/stores-ms/internal/infra/uow/interfaces"
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

func (u *UnitOfWork) Do(ctx context.Context, callback func(uow iuow.IUnitOfWork) error) error {
	if u.TX != nil {
		return fmt.Errorf("transaction is already open")
	}

	tx, err := u.DB.Begin(ctx)
	if err != nil {
		return err
	}

	u.TX = tx
	err = callback(u)
	if err != nil {
		errRb := u.Rollback()
		if errRb != nil {
			return errors.New(fmt.Sprintf("original error: %s, rollback error: %s", err.Error(), errRb.Error()))
		}
		return err
	}
	return u.CommitOrRollback()
}

func (u *UnitOfWork) Rollback() error {
	if u.TX == nil {
		return fmt.Errorf("no transaction to rollback")
	}

	err := u.TX.Rollback(context.Background())
	if err != nil {
		return err
	}

	u.TX = nil
	return nil
}

func (u *UnitOfWork) CommitOrRollback() error {
	if u.TX == nil {
		return nil
	}

	if err := u.TX.Commit(context.Background()); err != nil {
		if errRB := u.Rollback(); errRB != nil {
			return fmt.Errorf("erro original: %s, erro de rollback: %s", err.Error(), errRB.Error())
		}
		return err
	}

	u.TX = nil
	return nil
}
