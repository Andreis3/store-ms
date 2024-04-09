package igroup_controller

import "github.com/andreis3/stores-ms/internal/util"

type IGroupRouter interface {
	GroupRoutes() util.RouterType
}
