package uow

import (
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/andreis3/stores-ms/internal/infra/adapters/database/postgres/interfaces"
	"github.com/andreis3/stores-ms/internal/infra/repository/postgres/group"
	"github.com/andreis3/stores-ms/internal/util"
)

func NewProxyUnitOfWork(pool *pgxpool.Pool) *UnitOfWork {
	uow := NewUnitOfWork(pool)
	uow.Register(util.GROUP_REPOSITORY_KEY, func(tx any) any {
		repo := repo_group.NewGroupRepository()
		repo.DB = ipostgres.New(tx.(pgx.Tx))
		return repo
	})
	return uow
}
