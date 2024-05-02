package make_controller

import (
	"github.com/andreis3/stores-ms/internal/app/command/group"
	"github.com/andreis3/stores-ms/internal/domain/service/group"
	"github.com/andreis3/stores-ms/internal/infra/common/logger"
	"github.com/andreis3/stores-ms/internal/infra/common/metrics/prometheus"
	"github.com/andreis3/stores-ms/internal/infra/common/uuid"
	"github.com/andreis3/stores-ms/internal/infra/uow"
	"github.com/andreis3/stores-ms/internal/interfaces/http/controllers/group"
	"github.com/andreis3/stores-ms/internal/interfaces/http/controllers/group/interfaces"
	"github.com/jackc/pgx/v5/pgxpool"
)

func MakeSearchGroupController(pool *pgxpool.Pool) igroup_controller.ISearchGroupController {
	prometheus := metric_prometheus.NewPrometheusAdapter()
	unitOfWork := uow.NewProxyUnitOfWork(pool, prometheus)
	log := logger.NewLogger()
	searchGroupService := group_service.NewSearchGroupService(unitOfWork)
	uuid := uuid.NewUUID()
	searchGroupCommand := group_command.NewSearchGroupCommand(searchGroupService)
	searchGroupController := group_controller.NewSearchGroupController(searchGroupCommand, prometheus, log, uuid)
	return searchGroupController
}
