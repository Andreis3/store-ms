package make_controller

import (
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/andreis3/stores-ms/internal/app/command/group"
	"github.com/andreis3/stores-ms/internal/domain/service/group_service"
	"github.com/andreis3/stores-ms/internal/infra/common/logger"
	"github.com/andreis3/stores-ms/internal/infra/uow"
	"github.com/andreis3/stores-ms/internal/interface/http/group"
	"github.com/andreis3/stores-ms/internal/interface/http/group/interfaces"
)

func MakeControllerGroup(pool *pgxpool.Pool) igroup_controller.IGroupController {
	uow := uow.NewProxyUnitOfWork(pool)
	logger := logger.NewLogger()
	groupService := group_service.NewInsertGroupService(uow)
	groupCommand := group_command.NewInsertGroupCommand(groupService)
	groupController := group_controller.NewGroupController(groupCommand, logger)
	return groupController
}
