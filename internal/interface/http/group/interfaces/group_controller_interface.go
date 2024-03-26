package group_controller

import "net/http"

type IGroupController interface {
	CreateGroup(w http.ResponseWriter, r *http.Request)
}
