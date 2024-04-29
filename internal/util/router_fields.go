package util

import "net/http"

type RouterType []RouterFields
type RouterFields struct {
	Method      string
	Path        string
	Controller  any
	Description string
	Type        string
	Middlewares []func(http.Handler) http.Handler
}
