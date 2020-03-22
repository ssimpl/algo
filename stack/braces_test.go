package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBraces(t *testing.T) {
	t.Run("one operand", func(t *testing.T) {
		a := "(a)"
		out := Braces(a)
		assert.True(t, out)
	})

	t.Run("has redundant braces", func(t *testing.T) {
		a := "((a + b))"
		out := Braces(a)
		assert.True(t, out)
	})

	t.Run("doesn't have have redundant braces", func(t *testing.T) {
		a := "(a + (a + b))"
		out := Braces(a)
		assert.False(t, out)
	})
}
