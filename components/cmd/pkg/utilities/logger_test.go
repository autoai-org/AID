package utilities

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestNewLogger(t *testing.T) {
	logger := NewLogger()
	assert.Equal(t, os.Stdout, logger.Out)
}
