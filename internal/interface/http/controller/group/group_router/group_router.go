package group_router

import (
	"net/http"

	"github.com/andreis3/stores-ms/internal/interface/http/helpers"
	"github.com/andreis3/stores-ms/internal/util"

	"github.com/andreis3/stores-ms/internal/interface/http/controller/group/interfaces"
)

type GroupRouter struct {
	createGroupController igroup_controller.ICreateGroupController
}

func NewGroupRouter(createGroupController igroup_controller.ICreateGroupController) *GroupRouter {
	return &GroupRouter{
		createGroupController: createGroupController,
	}
}

func (r *GroupRouter) GroupRoutes() util.RouterType {
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
