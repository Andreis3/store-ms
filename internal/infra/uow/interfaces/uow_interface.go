package iuow

import (
	"context"

	"github.com/andreis3/stores-ms/internal/util"
)

type RepositoryFactory func(tx any) any

type IUnitOfWork interface {
	Register(name string, callback RepositoryFactory)
	GetRepository(name string) any
	Do(ctx context.Context, callback func(uow IUnitOfWork) *util.ValidationError) *util.ValidationError
	CommitOrRollback() *util.ValidationError
	Rollback() *util.ValidationError
}
