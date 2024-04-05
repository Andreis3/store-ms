package group_controller

import (
	"context"
	"github.com/andreis3/stores-ms/internal/infra/common/metrics/prometheus"
	"net/http"
	"strings"
	"time"

	"github.com/andreis3/stores-ms/internal/app/command/group/interfaces"
	"github.com/andreis3/stores-ms/internal/infra/common/logger/interfaces"
	"github.com/andreis3/stores-ms/internal/interface/http/controller/group/dto"
	"github.com/andreis3/stores-ms/internal/interface/http/helpers"
)

type Controller struct {
	insertGroupCommand igroup_command.IInsertGroupCommand
	logger             ilogger.ILogger
	requestID          helpers.IRequestID
	prometheus         prometheus.PrometheusAdapter
}

func NewGroupController(insertGroupCommand igroup_command.IInsertGroupCommand,
	logger ilogger.ILogger,
	requestID helpers.IRequestID) *Controller {
	prom := prometheus.NewPrometheusAdapter()
	return &Controller{
		insertGroupCommand: insertGroupCommand,
		logger:             logger,
		requestID:          requestID,
		prometheus:         *prom,
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
			"ERROR_MESSAGE", strings.Join(err.LogError, ", "))
		helpers.ResponseError[[]string](w, err.Status, requestID, err.Code, err.ClientError)
		c.prometheus.CounterRequestHttpStatusCode(context.Background(), err.Status)
		end := time.Now()
		duration := end.Sub(start).Milliseconds()
		c.prometheus.HistogramRequestDuration(context.Background(), "/groups", float64(duration))
		return
	}
	group, errCM := c.insertGroupCommand.Execute(*groupInputDTO)
	if errCM != nil {
		c.logger.Error("Create Group Error",
			"REQUEST_ID", requestID,
			"CODE_ERROR", errCM.Code,
			"ERROR_MESSAGE", strings.Join(errCM.LogError, ", "))
		helpers.ResponseError[[]string](w, errCM.Status, requestID, errCM.Code, errCM.ClientError)
		c.prometheus.CounterRequestHttpStatusCode(context.Background(), errCM.Status)
		end := time.Now()
		duration := end.Sub(start).Milliseconds()
		c.prometheus.HistogramRequestDuration(context.Background(), "/groups", float64(duration))
		return
	}
	c.prometheus.CounterRequestHttpStatusCode(context.Background(), http.StatusCreated)
	end := time.Now()
	duration := end.Sub(start).Milliseconds()
	c.prometheus.HistogramRequestDuration(context.Background(), "/groups", float64(duration))
	helpers.ResponseSuccess[group_dto.GroupOutputDTO](w, requestID, http.StatusCreated, group)
}
