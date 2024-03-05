package util

type NotificationContext struct {
	Notification []map[string]any
}

func (n *NotificationContext) AddNotification(notification map[string]any) {
	n.Notification = append(n.Notification, notification)
}

func (n *NotificationContext) HasNotification() bool {
	return len(n.Notification) > 0
}

func (n *NotificationContext) ReturnNotification() []map[string]any {
	return n.Notification
}
