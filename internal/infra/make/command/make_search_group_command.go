package make_command

import (
	"github.com/andreis3/stores-ms/internal/app/command"
	icommand "github.com/andreis3/stores-ms/internal/app/command/interfaces"
	"github.com/andreis3/stores-ms/internal/domain/services"
	"github.com/andreis3/stores-ms/internal/infra/adapters/database/interfaces"
	"github.com/andreis3/stores-ms/internal/infra/common/metrics/interface"
	"github.com/andreis3/stores-ms/internal/infra/uow"
	"github.com/jackc/pgx/v5/pgxpool"
)

func MakeSearchGroupCommand(db idatabase.IDatabase, prometheus imetric.IMetricAdapter) icommand.ISearchGroupCommand {
	pool := db.InstanceDB().(*pgxpool.Pool)
	unitOfWork := uow.NewProxyUnitOfWork(pool, prometheus)
	searchGroupService := services.NewSearchGroupService(unitOfWork)
	searchGroupCommand := command.NewSearchGroupCommand(searchGroupService)
	return searchGroupCommand
}
