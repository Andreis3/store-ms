package make_controller

import (
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/andreis3/stores-ms/internal/infra/common/metrics/prometheus"
	"github.com/andreis3/stores-ms/internal/infra/common/uuid"

	"github.com/andreis3/stores-ms/internal/app/command/group"
	"github.com/andreis3/stores-ms/internal/domain/service/group"
	"github.com/andreis3/stores-ms/internal/infra/common/logger"
	"github.com/andreis3/stores-ms/internal/infra/uow"
	"github.com/andreis3/stores-ms/internal/interface/http/controller/group"
	"github.com/andreis3/stores-ms/internal/interface/http/controller/group/interfaces"
)

func MakeCreateGroupController(pool *pgxpool.Pool) igroup_controller.ICreateGroupController {
	prometheus := metric_prometheus.NewPrometheusAdapter()
	uow := uow.NewProxyUnitOfWork(pool, prometheus)
	logger := logger.NewLogger()
	uuid := uuid.NewUUID()
	groupService := group_service.NewInsertGroupService(uow, uuid)
	groupCommand := group_command.NewInsertGroupCommand(groupService)
	groupController := group_controller.NewCreateGroupController(groupCommand, prometheus, logger, uuid)
	return groupController
}
