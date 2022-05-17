package p2

import (
	uuid "github.com/satori/go.uuid"
)

func New() uuid.UUID {
	return uuid.NewV4()
}
