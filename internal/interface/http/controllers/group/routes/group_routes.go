package group_routes

import (
	group_middleware "github.com/andreis3/stores-ms/internal/interface/http/controllers/group/middleware"
	"github.com/go-chi/chi/v5/middleware"
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

func (r *GroupRoutes) GroupRoutes() util.RouteType {
	return util.RouteType{
		{
			Method:      http.MethodPost,
			Path:        helpers.CREATE_GROUP_V1,
			Controller:  r.createGroupController.CreateGroup,
			Description: "Create Group",
			Type:        util.HANDLER_FUNC,
			Middlewares: []func(http.Handler) http.Handler{
				group_middleware.ValidatePath,
				middleware.Logger,
			},
		},
	}
}
