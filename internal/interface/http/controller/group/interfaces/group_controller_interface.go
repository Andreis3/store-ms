package igroup_controller

import "net/http"

type IGroupController interface {
	CreateGroup(w http.ResponseWriter, r *http.Request)
	GetGroup(w http.ResponseWriter, r *http.Request)
}
