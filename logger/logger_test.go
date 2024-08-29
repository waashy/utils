package logger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLogger(t *testing.T) {
	t.Run("DEBUG", func(t *testing.T) {
		logger, err := NewLogger("DEBUG")
		t.Log("logger : ", logger)
		assert.Nil(t, err)
	})
}
