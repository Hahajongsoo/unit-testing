package ex1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsStringLong(t *testing.T) {
	var actual bool = IsStringLong("abc")
	assert.Equal(t, false, actual)
}
