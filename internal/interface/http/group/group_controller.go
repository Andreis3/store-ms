package group_controller

import (
	"net/http"

	"github.com/andreis3/stores-ms/internal/app/command/group/interfaces"
	group_dto "github.com/andreis3/stores-ms/internal/interface/http/group/dto"
	"github.com/andreis3/stores-ms/internal/util"
)

type Controller struct {
	groupCommand group_command.IInsertGroupCommand
}

func NewGroupController(groupCommand group_command.IInsertGroupCommand) *Controller {
	return &Controller{
		groupCommand: groupCommand,
	}
}

func (p *Controller) CreateGroup(w http.ResponseWriter, r *http.Request) {
	groupInputDTO, erro := util.RecoverBody[group_dto.GroupInputDTO](r)
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
