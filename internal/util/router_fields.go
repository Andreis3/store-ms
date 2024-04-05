package util

type RouterType []RouterFields
type RouterFields struct {
	Method      string
	Path        string
	Controller  any
	Description string
	Type        string
}
