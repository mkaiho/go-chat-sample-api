package adapter

import (
	"github.com/mkaiho/go-chat-sample-api/entity"
	"github.com/mkaiho/go-chat-sample-api/usecase/port"
)

var _ port.IDManager = (*IDManager)(nil)

func NewIDManager(generator func() string) *IDManager {
	return &IDManager{
		generator: generator,
	}
}

type IDManager struct {
	generator func() string
}

func (m *IDManager) Generate() entity.ID {
	return entity.ParseID(m.generator())
}
