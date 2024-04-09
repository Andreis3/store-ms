package make_controller

import (
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/andreis3/stores-ms/internal/infra/common/metrics/prometheus"

	"github.com/andreis3/stores-ms/internal/app/command/group"
	"github.com/andreis3/stores-ms/internal/domain/service/group"
	"github.com/andreis3/stores-ms/internal/infra/common/logger"
	"github.com/andreis3/stores-ms/internal/infra/uow"
	"github.com/andreis3/stores-ms/internal/interface/http/controller/group"
	"github.com/andreis3/stores-ms/internal/interface/http/controller/group/interfaces"
	"github.com/andreis3/stores-ms/internal/interface/http/helpers"
)

func MakeControllerGroup(pool *pgxpool.Pool) igroup_controller.IGroupController {
	uow := uow.NewProxyUnitOfWork(pool)
	logger := logger.NewLogger()
	groupService := group_service.NewInsertGroupService(uow)
	groupCommand := group_command.NewInsertGroupCommand(groupService)
	requestID := helpers.NewRequestID()
	prometheus := metric_prometheus.NewPrometheusAdapter()
	groupController := group_controller.NewGroupController(groupCommand, prometheus, logger, requestID)
	return groupController
}
