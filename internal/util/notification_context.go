package util

type NotificationContext struct {
	Notification []string
}

func (n *NotificationContext) AddNotification(notification string) {
	n.Notification = append(n.Notification, notification)
}

func (n *NotificationContext) HasNotification() bool {
	return len(n.Notification) > 0
}

func (n *NotificationContext) ReturnNotification() []string {
	return n.Notification
}
