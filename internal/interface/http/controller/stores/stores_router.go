package stores_controller

import (
	"github.com/andreis3/stores-ms/internal/util"
	"net/http"

	"github.com/andreis3/stores-ms/internal/interface/http/controller/stores/interfaces"
	"github.com/andreis3/stores-ms/internal/interface/http/controller/stores/middleware"
)

type Router struct {
	controller istores_controller.IStoreController
}

func NewStoresRouter(controller istores_controller.IStoreController) *Router {
	return &Router{
		controller: controller,
	}
}
func (r *Router) StoresRoutes() []map[string]any {
	return []map[string]any{
		{
			util.METHOD:      http.MethodPost,
			util.PATH:        "/stores",
			util.CONTROLLER:  r.controller.CreateStores,
			util.DESCRIPTION: "Create Stores",
			util.TYPE:        util.HANDLER_FUNC,
		},
		{
			util.METHOD:      http.MethodPut,
			util.PATH:        "/stores",
			util.CONTROLLER:  stores_middleware.ValidatePath(r.controller.UpdateStores),
			util.DESCRIPTION: "Update Stores",
			util.TYPE:        util.HANDLER_FUNC,
		},
		{
			util.METHOD:      http.MethodGet,
			util.PATH:        "/stores/{id}",
			util.CONTROLLER:  r.controller.ListStoresByID,
			util.DESCRIPTION: "List Stores by id",
			util.TYPE:        util.HANDLER_FUNC,
		},
	}
}
