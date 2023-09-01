package ex3

import (
	"testing"
)

func TestNoValidation(t *testing.T) {
	IsStringLong("abc")
	IsStringLong("abcdef")
}
