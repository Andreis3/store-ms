package istores_controller

import "github.com/andreis3/stores-ms/internal/util"

type IStoresRouter interface {
	StoresRoutes() util.RouteType
}
