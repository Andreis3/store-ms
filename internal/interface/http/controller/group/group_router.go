package group_controller

import (
	"github.com/andreis3/stores-ms/internal/util"
	"net/http"

	"github.com/andreis3/stores-ms/internal/interface/http/controller/group/interfaces"
)

type Router struct {
	controller igroup_controller.IGroupController
}

func NewGroupRouter(controller igroup_controller.IGroupController) *Router {
	return &Router{
		controller: controller,
	}
}

func (r *Router) GroupRoutes() []map[string]any {
	return []map[string]any{
		{
			util.METHOD:      http.MethodPost,
			util.PATH:        "/groups",
			util.CONTROLLER:  r.controller.CreateGroup,
			util.DESCRIPTION: "Create Group",
			util.TYPE:        util.HANDLER_FUNC,
		},
	}
}
