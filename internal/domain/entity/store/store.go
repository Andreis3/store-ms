package store

import (
	"fmt"
	"slices"

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
	Status      string
	CNPJ        string
	Domain      string
	GroupCOD    string
	Contacts    []Contact
	util.NotificationContext
}

type Contact struct {
	Name  string
	Email string
	Phone string
	Ramal string
}

func NewStore(storeKey, companyName, status, cnpj, domain, groupCOD string, contacts []Contact) *Store {
	return &Store{
		StoreKey:    storeKey,
		CompanyName: companyName,
		Status:      status,
		CNPJ:        cnpj,
		Domain:      domain,
		GroupCOD:    groupCOD,
		Contacts:    contacts,
	}
}

func (s *Store) Validate() []map[string]string {
	notifications := make([]map[string]string, 0)
	if s.StoreKey == "" {
		notifications = append(notifications, map[string]string{"store_key": "is required"})
	}
	if s.CompanyName == "" {
		notifications = append(notifications, map[string]string{"company_name": "is required"})
	}
	if s.Status == "" {
		notifications = append(notifications, map[string]string{"status": "is required"})
	}
	if s.Status != "" && !slices.Contains(STATUS, s.Status) {
		notifications = append(notifications, map[string]string{"status": "is invalid, valid values are active or inactive"})
	}
	if s.CNPJ == "" {
		notifications = append(notifications, map[string]string{"cnpj": "is required"})
	}
	if s.Domain == "" {
		notifications = append(notifications, map[string]string{"domain": "is required"})
	}
	if s.GroupCOD == "" {
		notifications = append(notifications, map[string]string{"group_code": "is required"})
	}
	if len(s.Contacts) < 1 {
		notifications = append(notifications, map[string]string{"contacts": "min 1 contact is required"})
	}

	for index, contact := range s.Contacts {
		if contact.Name == "" {
			key := fmt.Sprintf("contacts[%d].name", index)
			notifications = append(notifications, map[string]string{key: "is required"})
		}
		if contact.Email == "" {
			key := fmt.Sprintf("contacts[%d].email", index)
			notifications = append(notifications, map[string]string{key: "is required"})
		}
		if contact.Phone == "" {
			key := fmt.Sprintf("contacts[%d].phone", index)
			notifications = append(notifications, map[string]string{key: "is required"})
		}

	}
	return notifications
}
