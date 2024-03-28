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
	groupCommand group_command.IInsertGroupCommand
	logger       ilogger.ILogger
}

func NewGroupController(groupCommand group_command.IInsertGroupCommand, logger ilogger.ILogger) *Controller {
	return &Controller{
		groupCommand: groupCommand,
		logger:       logger,
	}
}

func (p *Controller) CreateGroup(w http.ResponseWriter, r *http.Request) {
	requestID := uuid.New().String()
	groupInputDTO, err := util.RecoverBody[*group_dto.GroupInputDTO](r)
	if err != nil {
		util.Response(w, http.StatusBadRequest, err.Error())
		return
	}

	group, erroCM := p.groupCommand.Execute(*groupInputDTO)
	if erroCM != nil {
		p.logger.Error("Create Group Error", "REQUEST_ID", requestID, "ERROR_MESSAGE", strings.Join(erroCM.LogError, ", "))
		util.ResponseBadRequestError[[]string](w, erroCM.Status, requestID, erroCM.ClientError)
		return
	}

	util.Response(w, http.StatusCreated, group)

}
