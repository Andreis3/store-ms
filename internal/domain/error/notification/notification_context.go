package notification

type NotificationContext struct {
	notification []string
}

func (n *NotificationContext) AddNotification(notification string) {
	n.notification = append(n.notification, notification)
}

func (n *NotificationContext) HasNotification() bool {
	return len(n.notification) > 0
}

func (n *NotificationContext) ReturnNotification() []string {
	return n.notification
}
