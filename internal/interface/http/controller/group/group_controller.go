package group_controller

import (
	"net/http"
	"strings"

	"github.com/andreis3/stores-ms/internal/app/command/group/interfaces"
	"github.com/andreis3/stores-ms/internal/infra/common/logger/interfaces"
	"github.com/andreis3/stores-ms/internal/interface/http/controller/group/dto"
	"github.com/andreis3/stores-ms/internal/interface/http/helpers"
)

type Controller struct {
	insertGroupCommand igroup_command.IInsertGroupCommand
	logger             ilogger.ILogger
	requestID          helpers.IRequestID
}

func NewGroupController(insertGroupCommand igroup_command.IInsertGroupCommand,
	logger ilogger.ILogger,
	requestID helpers.IRequestID) *Controller {
	return &Controller{
		insertGroupCommand: insertGroupCommand,
		logger:             logger,
		requestID:          requestID,
	}
}

func (c *Controller) CreateGroup(w http.ResponseWriter, r *http.Request) {
	requestID := c.requestID.Generate()
	groupInputDTO, err := helpers.DecoderBodyRequest[*group_dto.GroupInputDTO](r)
	if err != nil {
		c.logger.Error("Create Group Error",
			"REQUEST_ID", requestID,
			"CODE_ERROR", err.Code,
			"ERROR_MESSAGE", strings.Join(err.LogError, ", "))
		helpers.ResponseError[[]string](w, err.Status, requestID, err.Code, err.ClientError)
		return
	}
	group, errCM := c.insertGroupCommand.Execute(*groupInputDTO)
	if errCM != nil {
		c.logger.Error("Create Group Error",
			"REQUEST_ID", requestID,
			"CODE_ERROR", errCM.Code,
			"ERROR_MESSAGE", strings.Join(errCM.LogError, ", "))
		helpers.ResponseError[[]string](w, errCM.Status, requestID, errCM.Code, errCM.ClientError)
		return
	}
	helpers.ResponseSuccess[group_dto.GroupOutputDTO](w, requestID, http.StatusCreated, group)
}
