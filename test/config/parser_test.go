package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/waashy/utils/config/parser"
)

func TestConfigParser(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		type Config struct {
			Key string `json:"key"`
		}

		var config Config

		err := parser.ConfigParser(`../../test/config/test_config.json`, &config)
		assert.Nil(t, err)
		assert.Equal(t, config.Key, "value")
	})
}
