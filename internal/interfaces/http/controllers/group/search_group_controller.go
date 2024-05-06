package group_controller

import (
	"context"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/andreis3/stores-ms/internal/infra/common/metrics/interface"
	"github.com/andreis3/stores-ms/internal/infra/common/uuid"
	"github.com/andreis3/stores-ms/internal/infra/uow"

	"github.com/andreis3/stores-ms/internal/app/command/group/interfaces"
	"github.com/andreis3/stores-ms/internal/infra/common/logger/interfaces"
	"github.com/andreis3/stores-ms/internal/interfaces/http/controllers/group/dto"
	"github.com/andreis3/stores-ms/internal/interfaces/http/helpers"
)

type SearchGroupController struct {
	selectGroupCommand igroup_command.ISearchGroupCommand
	logger             ilogger.ILogger
	prometheus         imetric.IMetricAdapter
	requestID          uuid.IUUID
	mutex              sync.Mutex
	db                 *pgxpool.Pool
}

func NewSearchGroupController(
	selectGroupCommand igroup_command.ISearchGroupCommand,
	prometheus imetric.IMetricAdapter,
	logger ilogger.ILogger,
	requestID uuid.IUUID,
	db *pgxpool.Pool) *SearchGroupController {
	return &SearchGroupController{
		selectGroupCommand: selectGroupCommand,
		logger:             logger,
		prometheus:         prometheus,
		requestID:          requestID,
		db:                 db,
	}
}

func (ggc *SearchGroupController) SearchOneGroup(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	uow := uow.NewProxyUnitOfWork(ggc.db, ggc.prometheus)
	requestID := ggc.requestID.Generate()
	err := helpers.PathRouterValidate(r, helpers.ID)
	if err != nil {
		ggc.logger.Error("Get Group Error",
			"REQUEST_ID", requestID,
			"CODE_ERROR", err.Code,
			"ORIGIN", err.Origin,
			"ERROR_MESSAGE", strings.Join(err.LogError, ", "))
		ggc.prometheus.CounterRequestHttpStatusCode(context.Background(), helpers.SEARCH_GROUP_V1, err.Status)
		end := time.Now()
		duration := end.Sub(start).Milliseconds()
		ggc.prometheus.HistogramRequestDuration(context.Background(), helpers.SEARCH_GROUP_V1, err.Status, float64(duration))
		helpers.ResponseError[[]string](w, err.Status, requestID, err.Code, err.ClientError)
		return
	}
	id := r.PathValue("id")
	//ggc.mutex.Lock()
	group, err := ggc.selectGroupCommand.Execute(id, uow)
	//ggc.mutex.Unlock()
	if err != nil {
		ggc.logger.Error("Select One Group Error",
			"REQUEST_ID", requestID,
			"CODE_ERROR", err.Code,
			"ORIGIN", err.Origin,
			"ERROR_MESSAGE", strings.Join(err.LogError, ", "))
		ggc.prometheus.CounterRequestHttpStatusCode(context.Background(), helpers.SEARCH_GROUP_V1, err.Status)
		end := time.Now()
		duration := end.Sub(start).Milliseconds()
		ggc.prometheus.HistogramRequestDuration(context.Background(), helpers.SEARCH_GROUP_V1, err.Status, float64(duration))
		helpers.ResponseError[[]string](w, err.Status, requestID, err.Code, err.ClientError)
		return
	}
	ggc.prometheus.CounterRequestHttpStatusCode(context.Background(), helpers.SEARCH_GROUP_V1, http.StatusOK)
	end := time.Now()
	duration := end.Sub(start).Milliseconds()
	ggc.prometheus.HistogramRequestDuration(context.Background(), helpers.SEARCH_GROUP_V1, http.StatusOK, float64(duration))
	helpers.ResponseSuccess[group_dto.GroupOutputDTO](w, requestID, http.StatusOK, group)
}
