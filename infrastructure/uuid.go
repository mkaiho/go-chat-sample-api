package infrastructure

import (
	"github.com/google/uuid"
	"github.com/mkaiho/go-chat-sample-api/adapter"
)

func NewUUID() uuid.UUID {
	return uuid.New()
}

func NewUUIDManager() *adapter.IDManager {
	var generator = func() string {
		return NewUUID().String()
	}
	return adapter.NewIDManager(generator)
}
