package stores_controller

import (
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
			"method":      http.MethodPost,
			"path":        "/stores",
			"handler":     r.controller.CreateStores,
			"description": "Create Stores",
		},
		{
			"method":      http.MethodPut,
			"path":        "/stores",
			"handler":     stores_middleware.ValidatePath(r.controller.UpdateStores),
			"description": "Update Stores",
		},
		{
			"method":      http.MethodGet,
			"path":        "/stores/{id}",
			"handler":     r.controller.ListStoresByID,
			"description": "List Stores by id",
		},
	}
}
