package group_controller

import (
	"context"
	"net/http"
	"strings"
	"time"

	imetric "github.com/andreis3/stores-ms/internal/infra/common/metrics/interface"
	"github.com/andreis3/stores-ms/internal/infra/common/uuid"

	"github.com/andreis3/stores-ms/internal/app/command/group/interfaces"
	"github.com/andreis3/stores-ms/internal/infra/common/logger/interfaces"
	"github.com/andreis3/stores-ms/internal/interface/http/controller/group/dto"
	"github.com/andreis3/stores-ms/internal/interface/http/helpers"
)

type Controller struct {
	insertGroupCommand igroup_command.IInsertGroupCommand
	selectGroupCommand igroup_command.ISelectGroupCommand
	logger             ilogger.ILogger
	requestID          uuid.IUUID
	prometheus         imetric.IMetricAdapter
}

func NewGroupController(insertGroupCommand igroup_command.IInsertGroupCommand,
	prometheus imetric.IMetricAdapter,
	logger ilogger.ILogger,
	requestID uuid.IUUID) *Controller {
	return &Controller{
		insertGroupCommand: insertGroupCommand,
		logger:             logger,
		requestID:          requestID,
		prometheus:         prometheus,
	}
}

func (c *Controller) CreateGroup(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	requestID := c.requestID.Generate()
	groupInputDTO, err := helpers.DecoderBodyRequest[*group_dto.GroupInputDTO](r)
	if err != nil {
		c.logger.Error("Create Group Error",
			"REQUEST_ID", requestID,
			"CODE_ERROR", err.Code,
			"ORIGIN", err.Origin,
			"ERROR_MESSAGE", strings.Join(err.LogError, ", "))
		c.prometheus.CounterRequestHttpStatusCode(context.Background(), helpers.CREATE_GROUP_V1, err.Status)
		end := time.Now()
		duration := end.Sub(start).Milliseconds()
		c.prometheus.HistogramRequestDuration(context.Background(), helpers.CREATE_GROUP_V1, err.Status, float64(duration))
		helpers.ResponseError[[]string](w, err.Status, requestID, err.Code, err.ClientError)
		return
	}
	group, errCM := c.insertGroupCommand.Execute(*groupInputDTO)
	if errCM != nil {
		c.logger.Error("Create Group Error",
			"REQUEST_ID", requestID,
			"CODE_ERROR", errCM.Code,
			"ORIGIN", errCM.Origin,
			"ERROR_MESSAGE", strings.Join(errCM.LogError, ", "))
		c.prometheus.CounterRequestHttpStatusCode(context.Background(), helpers.CREATE_GROUP_V1, errCM.Status)
		end := time.Now()
		duration := end.Sub(start).Milliseconds()
		c.prometheus.HistogramRequestDuration(context.Background(), helpers.CREATE_GROUP_V1, errCM.Status, float64(duration))
		helpers.ResponseError[[]string](w, errCM.Status, requestID, errCM.Code, errCM.ClientError)
		return
	}
	c.prometheus.CounterRequestHttpStatusCode(context.Background(), helpers.CREATE_GROUP_V1, http.StatusCreated)
	end := time.Now()
	duration := end.Sub(start).Milliseconds()
	c.prometheus.HistogramRequestDuration(context.Background(), helpers.CREATE_GROUP_V1, http.StatusCreated, float64(duration))
	helpers.ResponseSuccess[group_dto.GroupOutputDTO](w, requestID, http.StatusCreated, group)
}

func (c *Controller) GetGroup(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	requestID := c.requestID.Generate()
	err := helpers.PathRouterValidate(r, helpers.ID)
	if err != nil {
		c.logger.Error("Get Group Error",
			"REQUEST_ID", requestID,
			"CODE_ERROR", err.Code,
			"ORIGIN", err.Origin,
			"ERROR_MESSAGE", strings.Join(err.LogError, ", "))
		c.prometheus.CounterRequestHttpStatusCode(context.Background(), helpers.GET_GROUP_V1, err.Status)
		end := time.Now()
		duration := end.Sub(start).Milliseconds()
		c.prometheus.HistogramRequestDuration(context.Background(), helpers.GET_GROUP_V1, err.Status, float64(duration))
		return
	}
	id := r.PathValue("id")
	group, err := c.selectGroupCommand.Execute(id)
	if err != nil {
		c.logger.Error("Select One Group Error",
			"REQUEST_ID", requestID,
			"CODE_ERROR", err.Code,
			"ORIGIN", err.Origin,
			"ERROR_MESSAGE", strings.Join(err.LogError, ", "))
		c.prometheus.CounterRequestHttpStatusCode(context.Background(), helpers.GET_GROUP_V1, err.Status)
		end := time.Now()
		duration := end.Sub(start).Milliseconds()
		c.prometheus.HistogramRequestDuration(context.Background(), helpers.GET_GROUP_V1, err.Status, float64(duration))
		helpers.ResponseError[[]string](w, err.Status, requestID, err.Code, err.ClientError)
		return
	}
	c.prometheus.CounterRequestHttpStatusCode(context.Background(), helpers.GET_GROUP_V1, http.StatusOK)
	end := time.Now()
	duration := end.Sub(start).Milliseconds()
	c.prometheus.HistogramRequestDuration(context.Background(), helpers.GET_GROUP_V1, http.StatusOK, float64(duration))
	helpers.ResponseSuccess[group_dto.GroupOutputDTO](w, requestID, http.StatusOK, group)
}
