package util

type ValidationError struct {
	Code        string
	Status      int
	ClientError []string
	LogError    []string
}

func (v *ValidationError) ExistError() bool {
	result := v != nil && (len(v.ClientError) > 0 || len(v.LogError) > 0)
	return result
}
