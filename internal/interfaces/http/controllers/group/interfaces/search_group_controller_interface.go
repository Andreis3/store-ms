package igroup_controller

import "net/http"

type ISearchGroupController interface {
	SearchOneGroup(w http.ResponseWriter, r *http.Request)
}
