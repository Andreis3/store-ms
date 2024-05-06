package group_routes

import (
	"net/http"

	"github.com/go-chi/chi/v5/middleware"

	"github.com/andreis3/stores-ms/internal/interfaces/http/controllers/group/middleware"

	"github.com/andreis3/stores-ms/internal/interfaces/http/helpers"
	"github.com/andreis3/stores-ms/internal/util"

	"github.com/andreis3/stores-ms/internal/interfaces/http/controllers/group/interfaces"
)

type GroupRoutes struct {
	createGroupController igroup_controller.ICreateGroupController
	searchGroupController igroup_controller.ISearchGroupController
}

func NewGroupRoutes(
	createGroupController igroup_controller.ICreateGroupController,
	searchGroupController igroup_controller.ISearchGroupController) *GroupRoutes {
	return &GroupRoutes{
		createGroupController: createGroupController,
		searchGroupController: searchGroupController,
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
		{
			Method:      http.MethodGet,
			Path:        helpers.SEARCH_GROUP_V1,
			Controller:  r.searchGroupController.SearchOneGroup,
			Description: "Search Group",
			Type:        util.HANDLER_FUNC,
			Middlewares: []func(http.Handler) http.Handler{
				//group_middleware.ValidatePath,
				//middleware.Logger,
			},
		},
	}
}
