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

func (r *Router) GroupRoutes() util.RouterType {
	return util.RouterType{
		{
			Method:      http.MethodPost,
			Path:        "/groups",
			Controller:  r.controller.CreateGroup,
			Description: "Create Group",
			Type:        util.HANDLER_FUNC,
		},
	}
}
