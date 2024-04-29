package stores_controller

import (
	"net/http"

	"github.com/andreis3/stores-ms/internal/util"

	"github.com/andreis3/stores-ms/internal/interface/http/controllers/stores/interfaces"
)

type Router struct {
	controller istores_controller.IStoreController
}

func NewStoresRouter(controller istores_controller.IStoreController) *Router {
	return &Router{
		controller: controller,
	}
}
func (r *Router) StoresRoutes() util.RouteType {
	return util.RouteType{
		{
			Method:      http.MethodPost,
			Path:        "/api/v1/stores",
			Controller:  r.controller.CreateStores,
			Description: "Create Stores",
			Type:        util.HANDLER_FUNC,
		},
	}
}
