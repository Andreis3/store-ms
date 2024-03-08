package group

import (
	"github.com/andreis3/stores-ms/internal/domain/entity/base"
	"github.com/andreis3/stores-ms/internal/domain/valueobject"
	"github.com/andreis3/stores-ms/internal/util"
)

type Group struct {
	ID        string
	GroupName string
	Code      string
	Status    valueobject.Status
	util.NotificationContext
}

func NewGroup(groupName, code string, status *valueobject.Status) *Group {

	return &Group{
		ID:        base.NewID(),
		GroupName: groupName,
		Code:      code,
		Status:    *status,
	}
}

func (g *Group) Validate() []map[string]any {
	if g.GroupName == "" {
		g.AddNotification(map[string]any{"group_name": "is required"})
	}
	if g.Code == "" {
		g.AddNotification(map[string]any{"code": "is required"})
	}

	g.Status.Validate(&g.NotificationContext)

	return g.Notification
}
