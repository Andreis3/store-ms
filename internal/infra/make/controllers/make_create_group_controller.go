package make_controller

import (
	"github.com/andreis3/stores-ms/internal/app/command"
	"github.com/andreis3/stores-ms/internal/domain/services"
	"github.com/andreis3/stores-ms/internal/infra/adapters/database/interfaces"
	"github.com/andreis3/stores-ms/internal/infra/common/logger"
	"github.com/andreis3/stores-ms/internal/infra/common/metrics/interface"
	"github.com/andreis3/stores-ms/internal/infra/common/uuid"
	"github.com/andreis3/stores-ms/internal/infra/uow"
	"github.com/andreis3/stores-ms/internal/interfaces/http/controllers/group"
	"github.com/andreis3/stores-ms/internal/interfaces/http/controllers/group/interfaces"
	"github.com/jackc/pgx/v5/pgxpool"
)

func MakeCreateGroupController(db idatabase.IDatabase, prometheus imetric.IMetricAdapter) igroup_controller.ICreateGroupController {
	pool := db.InstanceDB().(*pgxpool.Pool)
	uow := uow.NewProxyUnitOfWork(pool, prometheus)
	logger := logger.NewLogger()
	uuid := uuid.NewUUID()
	createGroupService := services.NewCreateGroupService(uow, uuid)
	createGroupCommand := command.NewCreateGroupCommand(createGroupService)
	createGroupController := group_controller.NewCreateGroupController(createGroupCommand, prometheus, logger, uuid)
	return createGroupController
}
