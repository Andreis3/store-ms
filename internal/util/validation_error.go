package util

type ValidationError struct {
	Code        string
	Status      int
	ClientError []string
	LogError    []string
}
