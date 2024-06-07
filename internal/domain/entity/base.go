package entity

import "github.com/google/uuid"

type Base struct {
	ID uuid.UUID
}

func NewID() string {
	return uuid.New().String()
}
