package entity

import (
	error2 "github.com/andreis3/stores-ms/internal/domain/notification"
	"github.com/andreis3/stores-ms/internal/domain/valueobject"
)

type SalesChannel struct {
	ID           string
	SalesChannel string
	Code         string
	StoreKey     string
	Status       valueobject.Status
	ConfigSalesChannel
	error2.NotificationError
}
type ConfigSalesChannel struct {
	AutomaticActive bool
}

func NewSalesChannel(salesChannel, code string, status *valueobject.Status, config bool) *SalesChannel {
	return &SalesChannel{
		ID:           "",
		SalesChannel: salesChannel,
		Code:         code,
		Status:       *status,
		ConfigSalesChannel: ConfigSalesChannel{
			AutomaticActive: config,
		},
	}
}

func (sc *SalesChannel) Validate() error2.NotificationError {
	if sc.SalesChannel == "" {
		sc.AddNotification(`sales_channel: is required`)
	}
	if sc.Code == "" {
		sc.AddNotification(`code: is required`)
	}
	sc.Status.Validate(&sc.NotificationError)
	return sc.NotificationError
}
