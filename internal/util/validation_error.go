package util

type ValidationError struct {
	Status      int
	ClientError []string
	LogError    []string
}

func NewValidationError(erros []string, status int) *ValidationError {
	return &ValidationError{
		ClientError: erros,
		Status:      status,
	}
}

func (v *ValidationError) ExistError() bool {
	result := v != nil && (len(v.ClientError) > 0 || len(v.LogError) > 0)
	return result
}
