package entity

import (
	error2 "github.com/andreis3/stores-ms/internal/domain/notification"
	"github.com/andreis3/stores-ms/internal/domain/valueobject"
)

type Seller struct {
	ID         string
	SellerName string
	Code       string
	StoreKey   string
	Status     valueobject.Status
	ConfigSeller
	error2.NotificationError
}
type ConfigSeller struct {
	AutomaticActive bool
}

func NewSeller(sellerName, code string, status *valueobject.Status, config bool) *Seller {
	return &Seller{
		ID:         "",
		SellerName: sellerName,
		Code:       code,
		Status:     *status,
		ConfigSeller: ConfigSeller{
			AutomaticActive: config,
		},
	}
}
func (s *Seller) Validate() error2.NotificationError {
	if s.SellerName == "" {
		s.AddNotification(`seller_name: is required`)
	}
	if s.Code == "" {
		s.AddNotification(`code: is required`)
	}
	s.Status.Validate(&s.NotificationError)
	return s.NotificationError
}
