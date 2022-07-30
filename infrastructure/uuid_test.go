package infrastructure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUUIDManager(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "generated uuid is not empty",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewUUIDManager()
			if assert.NotNil(t, got, "NewUUIDManager() = %v", got) {
				id := got.Generate().String()
				assert.True(t, len(id) > 0)
			}
		})
	}
}
