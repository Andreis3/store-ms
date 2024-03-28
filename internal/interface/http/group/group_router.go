package group_controller

import (
	"net/http"

	"github.com/andreis3/stores-ms/internal/interface/http/group/interfaces"
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
			"method":      http.MethodPost,
			"path":        "/groups",
			"handler":     r.controller.CreateGroup,
			"description": "Create Group",
		},
	}
}
