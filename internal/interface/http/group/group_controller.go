package group_controller

import (
	"encoding/json"
	"fmt"
	"net/http"

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
	groupInputDTO, erro := util.RecoverBody[*group_dto.GroupInputDTO](r)
	data, _ := json.Marshal(groupInputDTO)
	p.logger.Debug("CreateGroup", fmt.Sprintf("%s", string(data)))
	if erro != nil {
		util.Response(w, http.StatusBadRequest, erro.Error())
		return
	}

	group, erroCM := p.groupCommand.Execute(*groupInputDTO)
	if erroCM != nil {
		util.ResponseBadRequestError(w, http.StatusBadRequest, erroCM.Error())
		return
	}

	util.Response(w, http.StatusCreated, group)

}
