package igroup_controller

import "net/http"

type ISearchGroupController interface {
	SearchGroup(w http.ResponseWriter, r *http.Request)
}
