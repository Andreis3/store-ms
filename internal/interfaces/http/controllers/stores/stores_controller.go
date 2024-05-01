package stores_controller

import (
	"net/http"
)

type Controller struct {
	//logger     ilogger.ILogger
	//requestID  helpers.IRequestID
	//prometheus imetric.IMetricAdapter
}

func NewStoresController() *Controller {
	return &Controller{}
}

func (p *Controller) CreateStores(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"id": "123"}`))
}
