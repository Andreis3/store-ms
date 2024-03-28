package stores

import (
	"net/http"

	stores "github.com/andreis3/stores-ms/internal/interface/http/stores/interfaces"
	"github.com/andreis3/stores-ms/internal/interface/http/stores/middleware"
)

type Router struct {
	controller stores.IStoreController
}

func NewStoresRouter(controller stores.IStoreController) *Router {
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
			"handler":     middleware.ValidatePath(r.controller.UpdateStores),
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
