package notification

type NotificationError struct {
	notification []string
}

func (n *NotificationError) AddNotification(notification string) {
	n.notification = append(n.notification, notification)
}
func (n *NotificationError) HasNotification() bool {
	return len(n.notification) > 0
}
func (n *NotificationError) ReturnNotification() []string {
	return n.notification
}
