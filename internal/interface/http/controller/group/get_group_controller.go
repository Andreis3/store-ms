package group_controller

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/andreis3/stores-ms/internal/infra/common/metrics/interface"
	"github.com/andreis3/stores-ms/internal/infra/common/uuid"

	"github.com/andreis3/stores-ms/internal/app/command/group/interfaces"
	"github.com/andreis3/stores-ms/internal/infra/common/logger/interfaces"
	"github.com/andreis3/stores-ms/internal/interface/http/controller/group/dto"
	"github.com/andreis3/stores-ms/internal/interface/http/helpers"
)

type GetGroupController struct {
	selectGroupCommand igroup_command.ISelectGroupCommand
	logger             ilogger.ILogger
	requestID          uuid.IUUID
	prometheus         imetric.IMetricAdapter
}

func NewGetGroupController(
	selectGroupCommand igroup_command.ISelectGroupCommand,
	prometheus imetric.IMetricAdapter,
	logger ilogger.ILogger,
	requestID uuid.IUUID) *GetGroupController {
	return &GetGroupController{
		selectGroupCommand: selectGroupCommand,
		logger:             logger,
		requestID:          requestID,
		prometheus:         prometheus,
	}
}

func (ggc *GetGroupController) GetGroup(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	requestID := ggc.requestID.Generate()
	err := helpers.PathRouterValidate(r, helpers.ID)
	if err != nil {
		ggc.logger.Error("Get Group Error",
			"REQUEST_ID", requestID,
			"CODE_ERROR", err.Code,
			"ORIGIN", err.Origin,
			"ERROR_MESSAGE", strings.Join(err.LogError, ", "))
		ggc.prometheus.CounterRequestHttpStatusCode(context.Background(), helpers.GET_GROUP_V1, err.Status)
		end := time.Now()
		duration := end.Sub(start).Milliseconds()
		ggc.prometheus.HistogramRequestDuration(context.Background(), helpers.GET_GROUP_V1, err.Status, float64(duration))
		helpers.ResponseError[[]string](w, err.Status, requestID, err.Code, err.ClientError)
		return
	}
	id := r.PathValue("id")
	group, err := ggc.selectGroupCommand.Execute(id)
	if err != nil {
		ggc.logger.Error("Select One Group Error",
			"REQUEST_ID", requestID,
			"CODE_ERROR", err.Code,
			"ORIGIN", err.Origin,
			"ERROR_MESSAGE", strings.Join(err.LogError, ", "))
		ggc.prometheus.CounterRequestHttpStatusCode(context.Background(), helpers.GET_GROUP_V1, err.Status)
		end := time.Now()
		duration := end.Sub(start).Milliseconds()
		ggc.prometheus.HistogramRequestDuration(context.Background(), helpers.GET_GROUP_V1, err.Status, float64(duration))
		helpers.ResponseError[[]string](w, err.Status, requestID, err.Code, err.ClientError)
		return
	}
	ggc.prometheus.CounterRequestHttpStatusCode(context.Background(), helpers.GET_GROUP_V1, http.StatusOK)
	end := time.Now()
	duration := end.Sub(start).Milliseconds()
	ggc.prometheus.HistogramRequestDuration(context.Background(), helpers.GET_GROUP_V1, http.StatusOK, float64(duration))
	helpers.ResponseSuccess[group_dto.GroupOutputDTO](w, requestID, http.StatusOK, group)
}
