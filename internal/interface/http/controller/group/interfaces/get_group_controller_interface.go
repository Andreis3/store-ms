package igroup_controller

import "net/http"

type IGetGroupController interface {
	GetGroup(w http.ResponseWriter, r *http.Request)
}
