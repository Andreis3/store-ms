package saleschannel

import (
	"slices"

	"github.com/andreis3/stores-ms/internal/util"
)

const (
	Active   = "active"
	Inactive = "inactive"
)

var STATUS = [...]string{Active, Inactive}

type SalesChannel struct {
	SalesChannel string
	Code         string
	Status       string
	StoreKey     string
	Config
	util.NotificationContext
}

type Config struct {
	AutomaticActive bool
}

func NewSalesChannel(salesChannel, code, status string, config bool) *SalesChannel {
	return &SalesChannel{
		SalesChannel: salesChannel,
		Code:         code,
		Status:       status,
		Config: Config{
			AutomaticActive: config,
		},
	}
}

func (sc *SalesChannel) Validate() []map[string]interface{} {
	if sc.SalesChannel == "" {
		sc.AddNotification(map[string]interface{}{"sales_channel": "is required"})
	}
	if sc.Code == "" {
		sc.AddNotification(map[string]interface{}{"code": "is required"})
	}
	if sc.Status == "" {
		sc.AddNotification(map[string]interface{}{"status": "is required"})
	}
	if sc.Status != "" && !slices.Contains(STATUS[:], sc.Status) {
		sc.AddNotification(map[string]interface{}{"status": "is invalid, valid values are active or inactive"})
	}
	return sc.Notification
}
