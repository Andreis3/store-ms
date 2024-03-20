package valueobject

import (
	"slices"

	"github.com/andreis3/stores-ms/internal/util"
)

const (
	ACTIVE   = "active"
	INACTIVE = "inactive"
)

var STATUS = [...]string{ACTIVE, INACTIVE}

type Status struct {
	Status string
}

func NewStatus(status string) *Status {
	return &Status{
		Status: status,
	}
}

func (s *Status) Validate(notification *util.NotificationContext) {
	if s.Status == "" {
		notification.AddNotification(map[string]any{"status": "is required"})
	}
	if s.Status != "" && !slices.Contains(STATUS[:], s.Status) {
		notification.AddNotification(map[string]any{"status": "is invalid, valid values are active or inactive"})
	}
}
