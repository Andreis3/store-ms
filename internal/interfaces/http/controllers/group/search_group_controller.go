package group_controller

import (
	"context"
	"github.com/andreis3/stores-ms/internal/infra/adapters/database/interfaces"
	"github.com/andreis3/stores-ms/internal/infra/make/command"
	"net/http"
	"strings"
	"time"

	"github.com/andreis3/stores-ms/internal/infra/common/metrics/interface"
	"github.com/andreis3/stores-ms/internal/infra/common/uuid"

	"github.com/andreis3/stores-ms/internal/infra/common/logger/interfaces"
	"github.com/andreis3/stores-ms/internal/interfaces/http/controllers/group/dto"
	"github.com/andreis3/stores-ms/internal/interfaces/http/helpers"
)

type SearchGroupController struct {
	logger     ilogger.ILogger
	prometheus imetric.IMetricAdapter
	requestID  uuid.IUUID
	db         idatabase.IDatabase
}

func NewSearchGroupController(
	db idatabase.IDatabase,
	prometheus imetric.IMetricAdapter,
	logger ilogger.ILogger,
	requestID uuid.IUUID) *SearchGroupController {
	return &SearchGroupController{
		logger:     logger,
		prometheus: prometheus,
		requestID:  requestID,
		db:         db,
	}
}

func (ggc *SearchGroupController) SearchOneGroup(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	selectGroupCommand := make_command.MakeSearchGroupCommand(ggc.db, ggc.prometheus)
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
	group, err := selectGroupCommand.Execute(id)
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
