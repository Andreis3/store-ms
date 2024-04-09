package stores_controller

import (
	"fmt"
	"net/http"
)

type Controller struct{}

func NewStoresController() *Controller {
	return &Controller{}
}

func (p *Controller) CreateStores(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"id": "123"}`))
}
func (p *Controller) UpdateStores(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`{"id": "%s"}`, id)))
}
func (p *Controller) ListStoresByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`{"id": "%s"}`, id)))
}
