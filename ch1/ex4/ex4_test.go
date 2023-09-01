package ex4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPars(t *testing.T) {
	actual := Parse("5")
	assert.Equal(t, 5, actual)
}
