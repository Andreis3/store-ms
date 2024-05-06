package entity_group

import (
	"github.com/andreis3/stores-ms/internal/domain/error/notification"
	"github.com/andreis3/stores-ms/internal/domain/valueobject"
)

type Group struct {
	ID     string
	Name   string
	Code   string
	Status valueobject.Status
	notification.NotificationError
}

func NewGroup(name, code string, status *valueobject.Status) *Group {
	return &Group{
		ID:     "",
		Name:   name,
		Code:   code,
		Status: *status,
	}
}
func (g *Group) Validate() *notification.NotificationError {
	if g.Name == "" {
		g.AddNotification(`name: is required`)
	}
	if g.Code == "" {
		g.AddNotification(`code: is required`)
	}
	g.Status.Validate(&g.NotificationError)
	return &g.NotificationError
}
