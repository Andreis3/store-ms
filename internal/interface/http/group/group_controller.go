package group_controller

import (
	"net/http"
	"strings"

	"github.com/google/uuid"

	"github.com/andreis3/stores-ms/internal/app/command/group/interfaces"
	"github.com/andreis3/stores-ms/internal/infra/common/logger/interfaces"
	"github.com/andreis3/stores-ms/internal/interface/http/group/dto"
	"github.com/andreis3/stores-ms/internal/util"
)

type Controller struct {
	insertGroupCommand igroup_command.IInsertGroupCommand
	logger             ilogger.ILogger
}

func NewGroupController(insertGroupCommand igroup_command.IInsertGroupCommand, logger ilogger.ILogger) *Controller {
	return &Controller{
		insertGroupCommand: insertGroupCommand,
		logger:             logger,
	}
}

func (c *Controller) CreateGroup(w http.ResponseWriter, r *http.Request) {
	requestID := uuid.New().String()
	groupInputDTO, err := util.DecoderBodyRequest[*group_dto.GroupInputDTO](r)
	if err != nil {
		c.logger.Error("Create Group Error",
			"REQUEST_ID", requestID,
			"CODE_ERROR", err.Code,
			"ERROR_MESSAGE", strings.Join(err.LogError, ", "))
		util.ResponseError[[]string](w, err.Status, requestID, err.Code, err.ClientError)
		return
	}
	group, errCM := c.insertGroupCommand.Execute(*groupInputDTO)
	if errCM != nil {
		c.logger.Error("Create Group Error",
			"REQUEST_ID", requestID,
			"CODE_ERROR", errCM.Code,
			"ERROR_MESSAGE", strings.Join(errCM.LogError, ", "))
		util.ResponseError[[]string](w, errCM.Status, requestID, errCM.Code, errCM.ClientError)
		return
	}
	util.ResponseSuccess[group_dto.GroupOutputDTO](w, requestID, http.StatusCreated, group)
}
