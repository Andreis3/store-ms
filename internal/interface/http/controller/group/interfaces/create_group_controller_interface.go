package igroup_controller

import "net/http"

type ICreateGroupController interface {
	CreateGroup(w http.ResponseWriter, r *http.Request)
}
