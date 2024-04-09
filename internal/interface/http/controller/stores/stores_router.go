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
func (r *Router) StoresRoutes() util.RouterType {
	return util.RouterType{
		{
			Method:      http.MethodPost,
			Path:        "/stores",
			Controller:  r.controller.CreateStores,
			Description: "Create Stores",
			Type:        util.HANDLER_FUNC,
		},
		{
			Method:      http.MethodPut,
			Path:        "/stores",
			Controller:  stores_middleware.ValidatePath(r.controller.UpdateStores),
			Description: "Update Stores",
			Type:        util.HANDLER_FUNC,
		},
		{
			Method:      http.MethodGet,
			Path:        "/stores/{id}",
			Controller:  r.controller.ListStoresByID,
			Description: "List Stores by id",
			Type:        util.HANDLER_FUNC,
		},
	}
}
