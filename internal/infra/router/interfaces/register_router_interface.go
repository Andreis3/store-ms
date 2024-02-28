package interfaces

import "net/http"

type IRegisterRouter interface {
	Register(app *http.ServeMux, router []map[string]any)
}
