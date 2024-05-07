package make_command

import (
	"github.com/andreis3/stores-ms/internal/app/command/group"
	"github.com/andreis3/stores-ms/internal/app/command/group/interfaces"
	"github.com/andreis3/stores-ms/internal/domain/service/group"
	"github.com/andreis3/stores-ms/internal/infra/adapters/database/interfaces"
	"github.com/andreis3/stores-ms/internal/infra/common/metrics/interface"
	"github.com/andreis3/stores-ms/internal/infra/uow"
	"github.com/jackc/pgx/v5/pgxpool"
)

func MakeSearchGroupCommand(db idatabase.IDatabase, prometheus imetric.IMetricAdapter) igroup_command.ISearchGroupCommand {
	pool := db.InstanceDB().(*pgxpool.Pool)
	unitOfWork := uow.NewProxyUnitOfWork(pool, prometheus)
	searchGroupService := group_service.NewSearchGroupService(unitOfWork)
	searchGroupCommand := group_command.NewSearchGroupCommand(searchGroupService)
	return searchGroupCommand
}
