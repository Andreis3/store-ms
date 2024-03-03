package util

type NotificationContext struct {
	Notification []map[string]interface{}
}

func (n *NotificationContext) AddNotification(notification map[string]interface{}) {
	n.Notification = append(n.Notification, notification)
}

func (n *NotificationContext) HasNotification() bool {
	return len(n.Notification) > 0
}

func (n *NotificationContext) ReturnNotification() []map[string]interface{} {
	return n.Notification
}
