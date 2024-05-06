package make_controller

import (
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/andreis3/stores-ms/internal/infra/common/logger"
	imetric "github.com/andreis3/stores-ms/internal/infra/common/metrics/interface"
	"github.com/andreis3/stores-ms/internal/infra/common/uuid"
	"github.com/andreis3/stores-ms/internal/interfaces/http/controllers/group"
	"github.com/andreis3/stores-ms/internal/interfaces/http/controllers/group/interfaces"
)

func MakeSearchGroupController(pool *pgxpool.Pool, prometheus imetric.IMetricAdapter) igroup_controller.ISearchGroupController {
	log := logger.NewLogger()
	requestID := uuid.NewUUID()
	searchGroupController := group_controller.NewSearchGroupController(pool, prometheus, log, requestID)
	return searchGroupController
}
