package stores_controller

import (
	"net/http"

	"github.com/andreis3/stores-ms/internal/util"

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
			Path:        "/api/v1/stores",
			Controller:  r.controller.CreateStores,
			Description: "Create Stores",
			Type:        util.HANDLER_FUNC,
		},
		{
			Method:      http.MethodPut,
			Path:        "/api/v1/stores",
			Controller:  stores_middleware.ValidatePath(r.controller.UpdateStores),
			Description: "Update Stores",
			Type:        util.HANDLER_FUNC,
		},
		{
			Method:      http.MethodGet,
			Path:        "/api/v1/stores/{id}",
			Controller:  r.controller.ListStoresByID,
			Description: "List Stores by id",
			Type:        util.HANDLER_FUNC,
		},
	}
}
