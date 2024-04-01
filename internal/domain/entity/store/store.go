package entity_store

import (
	"fmt"

	"github.com/andreis3/stores-ms/internal/domain/error/notification"
	"github.com/andreis3/stores-ms/internal/domain/valueobject"
)

type Store struct {
	StoreKey    string
	CompanyName string
	CNPJ        valueobject.CNPJ
	Domain      string
	GroupCOD    string
	Status      valueobject.Status
	Contacts    []Contact
	notification.NotificationContext
}
type Contact struct {
	Name  string
	Email string
	Phone string
	Ramal string
}

func NewStore(storeKey, companyName, domain, groupCOD string, cnpj *valueobject.CNPJ, status *valueobject.Status, contacts []Contact) *Store {
	return &Store{
		StoreKey:    storeKey,
		CompanyName: companyName,
		Status:      *status,
		CNPJ:        *cnpj,
		Domain:      domain,
		GroupCOD:    groupCOD,
		Contacts:    contacts,
	}
}

func (s *Store) Validate() notification.NotificationContext {
	if s.StoreKey == "" {
		s.AddNotification(`store_key: is required`)
	}
	if s.CompanyName == "" {
		s.AddNotification(`company_name: is required`)
	}
	s.Status.Validate(&s.NotificationContext)
	s.CNPJ.Validate(&s.NotificationContext)
	if s.Domain == "" {
		s.AddNotification(`domain: is required`)
	}
	if s.GroupCOD == "" {
		s.AddNotification(`code: is required`)
	}
	if len(s.Contacts) < 1 {
		s.AddNotification(`contacts: min 1 contact is required`)
	}
	for index, contact := range s.Contacts {
		if contact.Name == "" {
			key := fmt.Sprintf("contacts[%d].name", index)
			s.AddNotification(fmt.Sprintf(`%s: is required`, key))
		}
		if contact.Email == "" {
			key := fmt.Sprintf("contacts[%d].email", index)
			s.AddNotification(fmt.Sprintf(`%s: is required`, key))
		}
		if contact.Phone == "" {
			key := fmt.Sprintf("contacts[%d].phone", index)
			s.AddNotification(fmt.Sprintf(`%s: is required`, key))
		}
	}
	return s.NotificationContext
}
