package iuow

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type RepositoryFactory func(tx pgx.Tx) any

type IUnitOfWork interface {
	Register(name string, callback RepositoryFactory)
	GetRepository(ctx context.Context, name string) any
	Do(ctx context.Context, callback func(uow IUnitOfWork) error) error
	CommitOrRollback() error
	Rollback() error
}
