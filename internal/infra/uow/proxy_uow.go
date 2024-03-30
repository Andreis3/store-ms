package uow

import (
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/andreis3/stores-ms/internal/infra/repository/postgres/group"
	"github.com/andreis3/stores-ms/internal/util"
)

func NewProxyUnitOfWork(db *pgxpool.Pool) *UnitOfWork {
	uow := NewUnitOfWork(db)
	uow.Register(util.GROUP_REPOSITORY_KEY, func(tx pgx.Tx) any {
		return repo_group.NewGroupRepository(db)
	})
	return uow
}
