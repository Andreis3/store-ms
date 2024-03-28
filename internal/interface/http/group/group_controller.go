package group_controller

import (
	"net/http"
	"strings"

	"github.com/google/uuid"

	"github.com/andreis3/stores-ms/internal/app/command/group/interfaces"
	ilogger "github.com/andreis3/stores-ms/internal/infra/common/logger/interfaces"
	group_dto "github.com/andreis3/stores-ms/internal/interface/http/group/dto"
	"github.com/andreis3/stores-ms/internal/util"
)

type Controller struct {
	insertGroupCommand group_command.IInsertGroupCommand
	logger             ilogger.ILogger
}

func NewGroupController(groupCommand group_command.IInsertGroupCommand, logger ilogger.ILogger) *Controller {
	return &Controller{
		insertGroupCommand: groupCommand,
		logger:             logger,
	}
}

func (c *Controller) CreateGroup(w http.ResponseWriter, r *http.Request) {
	requestID := uuid.New().String()
	groupInputDTO, err := util.RecoverBody[*group_dto.GroupInputDTO](r)
	if err != nil {
		util.Response(w, http.StatusBadRequest, err.Error())
		return
	}

	group, errCM := c.insertGroupCommand.Execute(*groupInputDTO)
	if errCM != nil {
		c.logger.Error("Create Group Error", "REQUEST_ID", requestID, "ERROR_MESSAGE", strings.Join(errCM.LogError, ", "))
		util.ResponseBadRequestError[[]string](w, errCM.Status, requestID, errCM.ClientError)
		return
	}

	util.Response(w, http.StatusCreated, group)

}
