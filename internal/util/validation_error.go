package util

type ValidationError struct {
	Message      []string `json:"message"`
	notification []string
}

func NewValidationError(notification []string) *ValidationError {
	return &ValidationError{
		notification: notification,
	}
}

func (v *ValidationError) Error() []string {
	for _, value := range v.notification {
		v.Message = append(v.Message, value)
	}
	return v.Message
}
