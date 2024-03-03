package group

import (
	"slices"

	"github.com/andreis3/stores-ms/internal/util"
)

const (
	Active   = "active"
	Inactive = "inactive"
)

var STATUS = [...]string{Active, Inactive}

type Group struct {
	GroupName string
	Code      string
	Status    string
	util.NotificationContext
}

func NewGroup(groupName, code, status string) *Group {
	return &Group{
		GroupName: groupName,
		Code:      code,
		Status:    status,
	}
}

func (g *Group) Validate() []map[string]interface{} {
	if g.GroupName == "" {
		g.AddNotification(map[string]interface{}{"group_name": "is required"})
	}
	if g.Code == "" {
		g.AddNotification(map[string]interface{}{"code": "is required"})
	}
	if g.Status == "" {
		g.AddNotification(map[string]interface{}{"status": "is required"})
	}

	if g.Status != "" && !slices.Contains(STATUS[:], g.Status) {
		g.AddNotification(map[string]interface{}{"status": "is invalid, valid values are active or inactive"})
	}
	return g.Notification
}
