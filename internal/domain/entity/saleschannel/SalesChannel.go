package saleschannel

import (
	"github.com/andreis3/stores-ms/internal/domain/valueobject"
	"github.com/andreis3/stores-ms/internal/util"
)

type SalesChannel struct {
	SalesChannel string
	Code         string
	StoreKey     string
	Status       valueobject.Status
	Config
	util.NotificationContext
}

type Config struct {
	AutomaticActive bool
}

func NewSalesChannel(salesChannel, code string, status *valueobject.Status, config bool) *SalesChannel {
	return &SalesChannel{
		SalesChannel: salesChannel,
		Code:         code,
		Status:       *status,
		Config: Config{
			AutomaticActive: config,
		},
	}
}

func (sc *SalesChannel) Validate() []map[string]any {
	if sc.SalesChannel == "" {
		sc.AddNotification(map[string]any{"sales_channel": "is required"})
	}
	if sc.Code == "" {
		sc.AddNotification(map[string]any{"code": "is required"})
	}
	sc.Status.Validate(&sc.NotificationContext)

	return sc.Notification
}
