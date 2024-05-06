package uuid

import "github.com/google/uuid"

type IUUID interface {
	Generate() string
}

type RequestID struct{}

func NewUUID() *RequestID {
	return &RequestID{}
}

func (r *RequestID) Generate() string {
	return uuid.New().String()
}
