package group_routes

import (
	"net/http"

	"github.com/andreis3/stores-ms/internal/interface/http/helpers"
	"github.com/andreis3/stores-ms/internal/util"

	"github.com/andreis3/stores-ms/internal/interface/http/controllers/group/interfaces"
)

type GroupRoutes struct {
	createGroupController igroup_controller.ICreateGroupController
}

func NewGroupRoutes(createGroupController igroup_controller.ICreateGroupController) *GroupRoutes {
	return &GroupRoutes{
		createGroupController: createGroupController,
	}
}

func (r *GroupRoutes) GroupRoutes() util.RouterType {
	return util.RouterType{
		{
			Method:      http.MethodPost,
			Path:        helpers.CREATE_GROUP_V1,
			Controller:  r.createGroupController.CreateGroup,
			Description: "Create Group",
			Type:        util.HANDLER_FUNC,
		},
	}
}
