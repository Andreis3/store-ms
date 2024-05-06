package make_command

import (
	igroup_command "github.com/andreis3/stores-ms/internal/app/command/group/interfaces"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/andreis3/stores-ms/internal/app/command/group"
	"github.com/andreis3/stores-ms/internal/domain/service/group"
	imetric "github.com/andreis3/stores-ms/internal/infra/common/metrics/interface"
	"github.com/andreis3/stores-ms/internal/infra/uow"
)

func MakeSearchGroupCommand(pool *pgxpool.Pool, prometheus imetric.IMetricAdapter) igroup_command.ISearchGroupCommand {
	unitOfWork := uow.NewProxyUnitOfWork(pool, prometheus)
	searchGroupService := group_service.NewSearchGroupService(unitOfWork)
	searchGroupCommand := group_command.NewSearchGroupCommand(searchGroupService)
	return searchGroupCommand
}
