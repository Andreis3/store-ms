package entity_group

import (
	"github.com/andreis3/stores-ms/internal/domain/entity/base"
	"github.com/andreis3/stores-ms/internal/domain/error/notification"
	"github.com/andreis3/stores-ms/internal/domain/valueobject"
)

type Group struct {
	ID        string
	GroupName string
	Code      string
	Status    valueobject.Status
	notification.NotificationError
}

func NewGroup(groupName, code string, status *valueobject.Status) *Group {
	return &Group{
		ID:        base.NewID(),
		GroupName: groupName,
		Code:      code,
		Status:    *status,
	}
}
func (g *Group) Validate() *notification.NotificationError {
	if g.GroupName == "" {
		g.AddNotification(`group_name: is required`)
	}
	if g.Code == "" {
		g.AddNotification(`code: is required`)
	}
	g.Status.Validate(&g.NotificationError)
	return &g.NotificationError
}
