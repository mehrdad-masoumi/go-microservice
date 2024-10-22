package validator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidateEmail(t *testing.T) {
	assert.True(t, isValidEmail("test@gmail.com"))
	assert.False(t, isValidEmail("test@gmail"))
}
