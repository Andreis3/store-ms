package seller

import (
	"slices"

	"github.com/andreis3/stores-ms/internal/util"
)

const (
	Active   = "active"
	Inactive = "inactive"
)

var STATUS = [...]string{Active, Inactive}

type Seller struct {
	SellerName string
	Code       string
	Status     string
	StoreKey   string
	Config
	util.NotificationContext
}

type Config struct {
	AutomaticActive bool
}

func NewSeller(sellerName, code, status string, config bool) *Seller {
	return &Seller{
		SellerName: sellerName,
		Code:       code,
		Status:     status,
		Config: Config{
			AutomaticActive: config,
		},
	}
}

func (s *Seller) Validate() []map[string]interface{} {
	if s.SellerName == "" {
		s.AddNotification(map[string]interface{}{"seller_name": "is required"})
	}
	if s.Code == "" {
		s.AddNotification(map[string]interface{}{"code": "is required"})
	}
	if s.Status == "" {
		s.AddNotification(map[string]interface{}{"status": "is required"})
	}
	if s.Status != "" && !slices.Contains(STATUS[:], s.Status) {
		s.AddNotification(map[string]interface{}{"status": "is invalid, valid values are active or inactive"})
	}
	return s.Notification
}
