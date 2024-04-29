package istores_controller

import "net/http"

type IStoreController interface {
	CreateStores(w http.ResponseWriter, r *http.Request)
}
