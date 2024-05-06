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
	"github.com/andreis3/stores-ms/internal/interfaces/http/controllers/group/dto"
	"github.com/andreis3/stores-ms/internal/interfaces/http/helpers"
)

type CreateGroupController struct {
	createGroupCommand igroup_command.ICreateGroupCommand
	logger             ilogger.ILogger
	requestID          uuid.IUUID
	prometheus         imetric.IMetricAdapter
}

func NewCreateGroupController(
	createGroupCommand igroup_command.ICreateGroupCommand,
	prometheus imetric.IMetricAdapter,
	logger ilogger.ILogger,
	requestID uuid.IUUID) *CreateGroupController {
	return &CreateGroupController{
		createGroupCommand: createGroupCommand,
		logger:             logger,
		requestID:          requestID,
		prometheus:         prometheus,
	}
}

func (cgc *CreateGroupController) CreateGroup(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	requestID := cgc.requestID.Generate()
	groupInputDTO, err := helpers.DecoderBodyRequest[*group_dto.GroupInputDTO](r)
	if err != nil {
		cgc.logger.Error("Create Group Error",
			"REQUEST_ID", requestID,
			"CODE_ERROR", err.Code,
			"ORIGIN", err.Origin,
			"ERROR_MESSAGE", strings.Join(err.LogError, ", "))
		cgc.prometheus.CounterRequestHttpStatusCode(context.Background(), helpers.CREATE_GROUP_V1, err.Status)
		end := time.Now()
		duration := end.Sub(start).Milliseconds()
		cgc.prometheus.HistogramRequestDuration(context.Background(), helpers.CREATE_GROUP_V1, err.Status, float64(duration))
		helpers.ResponseError[[]string](w, err.Status, requestID, err.Code, err.ClientError)
		return
	}
	group, errCM := cgc.createGroupCommand.Execute(*groupInputDTO)
	if errCM != nil {
		cgc.logger.Error("Create Group Error",
			"REQUEST_ID", requestID,
			"CODE_ERROR", errCM.Code,
			"ORIGIN", errCM.Origin,
			"ERROR_MESSAGE", strings.Join(errCM.LogError, ", "))
		cgc.prometheus.CounterRequestHttpStatusCode(context.Background(), helpers.CREATE_GROUP_V1, errCM.Status)
		end := time.Now()
		duration := end.Sub(start).Milliseconds()
		cgc.prometheus.HistogramRequestDuration(context.Background(), helpers.CREATE_GROUP_V1, errCM.Status, float64(duration))
		helpers.ResponseError[[]string](w, errCM.Status, requestID, errCM.Code, errCM.ClientError)
		return
	}
	cgc.prometheus.CounterRequestHttpStatusCode(context.Background(), helpers.CREATE_GROUP_V1, http.StatusCreated)
	end := time.Now()
	duration := end.Sub(start).Milliseconds()
	cgc.prometheus.HistogramRequestDuration(context.Background(), helpers.CREATE_GROUP_V1, http.StatusCreated, float64(duration))
	helpers.ResponseSuccess[group_dto.GroupOutputDTO](w, requestID, http.StatusCreated, group)
}
