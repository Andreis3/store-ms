package make_controller

import (
	"github.com/andreis3/stores-ms/internal/infra/adapters/database/interfaces"
	"github.com/andreis3/stores-ms/internal/infra/common/logger"
	"github.com/andreis3/stores-ms/internal/infra/common/metrics/interface"
	"github.com/andreis3/stores-ms/internal/infra/common/uuid"
	"github.com/andreis3/stores-ms/internal/interfaces/http/controllers/group"
	"github.com/andreis3/stores-ms/internal/interfaces/http/controllers/group/interfaces"
)

func MakeSearchGroupController(db idatabase.IDatabase, prometheus imetric.IMetricAdapter) igroup_controller.ISearchGroupController {
	log := logger.NewLogger()
	requestID := uuid.NewUUID()
	searchGroupController := group_controller.NewSearchGroupController(db, prometheus, log, requestID)
	return searchGroupController
}
