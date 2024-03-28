package entity_seller

import (
	"github.com/andreis3/stores-ms/internal/domain/valueobject"
	"github.com/andreis3/stores-ms/internal/util"
)

type Seller struct {
	SellerName string
	Code       string
	StoreKey   string
	Status     valueobject.Status
	Config
	util.NotificationContext
}
type Config struct {
	AutomaticActive bool
}

func NewSeller(sellerName, code string, status *valueobject.Status, config bool) *Seller {
	return &Seller{
		SellerName: sellerName,
		Code:       code,
		Status:     *status,
		Config: Config{
			AutomaticActive: config,
		},
	}
}
func (s *Seller) Validate() []string {
	if s.SellerName == "" {
		s.AddNotification(`seller_name: is required`)
	}
	if s.Code == "" {
		s.AddNotification(`code: is required`)
	}
	s.Status.Validate(&s.NotificationContext)
	return s.Notification
}
