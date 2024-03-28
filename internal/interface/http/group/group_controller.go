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

func NewGroupController(groupCommand igroup_command.IInsertGroupCommand, logger ilogger.ILogger) *Controller {
	return &Controller{
		insertGroupCommand: groupCommand,
		logger:             logger,
	}
}
func (c *Controller) CreateGroup(w http.ResponseWriter, r *http.Request) {
	requestID := uuid.New().String()
	groupInputDTO, err := util.RecoverBody[*group_dto.GroupInputDTO](r)
	if err != nil {
		errValidate := util.ValidationError{
			LogError:    []string{err.Error()},
			ClientError: []string{"payload sent is invalid"},
			Status:      http.StatusBadRequest,
		}
		c.logger.Error("Create Group Error", "REQUEST_ID", requestID, "ERROR_MESSAGE", strings.Join(errValidate.LogError, ", "))
		util.ResponseError[[]string](w, errValidate.Status, requestID, errValidate.ClientError)
		return
	}
	group, errCM := c.insertGroupCommand.Execute(*groupInputDTO)
	if errCM != nil {
		c.logger.Error("Create Group Error", "REQUEST_ID", requestID, "ERROR_MESSAGE", strings.Join(errCM.LogError, ", "))
		util.ResponseError[[]string](w, errCM.Status, requestID, errCM.ClientError)
		return
	}
	util.ResponseSuccess[group_dto.GroupOutputDTO](w, requestID, http.StatusCreated, group)
}
