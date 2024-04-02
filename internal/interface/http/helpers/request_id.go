package helpers

import "github.com/google/uuid"

type IRequestID interface {
	Generate() string
}

type RequestID struct{}

func NewRequestID() *RequestID {
	return &RequestID{}
}

func (r *RequestID) Generate() string {
	return uuid.New().String()
}
