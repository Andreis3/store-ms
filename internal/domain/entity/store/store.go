package entity_store

import (
	"fmt"

	"github.com/andreis3/stores-ms/internal/domain/valueobject"
	"github.com/andreis3/stores-ms/internal/util"
)

const (
	Active   = "active"
	Inactive = "inactive"
)

var STATUS = []string{Active, Inactive}

type Store struct {
	StoreKey    string
	CompanyName string
	CNPJ        valueobject.CNPJ
	Domain      string
	GroupCOD    string
	Status      valueobject.Status
	Contacts    []Contact
	util.NotificationContext
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

func (s *Store) Validate() []map[string]any {
	if s.StoreKey == "" {
		s.AddNotification(map[string]any{"store_key": "is required"})
	}
	if s.CompanyName == "" {
		s.AddNotification(map[string]any{"company_name": "is required"})
	}

	s.Status.Validate(&s.NotificationContext)

	s.CNPJ.Validate(&s.NotificationContext)

	if s.Domain == "" {
		s.AddNotification(map[string]any{"domain": "is required"})
	}
	if s.GroupCOD == "" {
		s.AddNotification(map[string]any{"group_code": "is required"})
	}
	if len(s.Contacts) < 1 {
		s.AddNotification(map[string]any{"contacts": "min 1 contact is required"})
	}

	for index, contact := range s.Contacts {
		if contact.Name == "" {
			key := fmt.Sprintf("contacts[%d].name", index)
			s.AddNotification(map[string]any{key: "is required"})
		}
		if contact.Email == "" {
			key := fmt.Sprintf("contacts[%d].email", index)
			s.AddNotification(map[string]any{key: "is required"})
		}
		if contact.Phone == "" {
			key := fmt.Sprintf("contacts[%d].phone", index)
			s.AddNotification(map[string]any{key: "is required"})
		}

	}
	return s.Notification
}
