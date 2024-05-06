package make_controller

import (
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/andreis3/stores-ms/internal/app/command/group"
	"github.com/andreis3/stores-ms/internal/domain/service/group"
	"github.com/andreis3/stores-ms/internal/infra/common/logger"
	imetric "github.com/andreis3/stores-ms/internal/infra/common/metrics/interface"
	"github.com/andreis3/stores-ms/internal/infra/common/uuid"
	"github.com/andreis3/stores-ms/internal/infra/uow"
	"github.com/andreis3/stores-ms/internal/interfaces/http/controllers/group"
	"github.com/andreis3/stores-ms/internal/interfaces/http/controllers/group/interfaces"
)

func MakeSearchGroupController(pool *pgxpool.Pool, prometheus imetric.IMetricAdapter) igroup_controller.ISearchGroupController {
	unitOfWork := uow.NewProxyUnitOfWork(pool, prometheus)
	log := logger.NewLogger()
	searchGroupService := group_service.NewSearchGroupService(unitOfWork)
	requestID := uuid.NewUUID()
	searchGroupCommand := group_command.NewSearchGroupCommand(searchGroupService)
	searchGroupController := group_controller.NewSearchGroupController(searchGroupCommand, prometheus, log, requestID)
	return searchGroupController
}
