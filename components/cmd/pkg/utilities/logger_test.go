package utilities

import (
	"os"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNewLogger(t *testing.T) {
	logger := NewLogger()
	assert.Equal(t, os.Stdout, logger.Out)
}
