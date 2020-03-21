package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParenthesesIsValid(t *testing.T) {
	t.Run("correct", func(t *testing.T) {
		a := "()[]{}"
		out := ParenthesesIsValid(a)
		assert.True(t, out)
	})

	t.Run("correct internal", func(t *testing.T) {
		a := "([{}]{{}})"
		out := ParenthesesIsValid(a)
		assert.True(t, out)
	})

	t.Run("invalid order", func(t *testing.T) {
		a := "([)]"
		out := ParenthesesIsValid(a)
		assert.False(t, out)
	})

	t.Run("closed first", func(t *testing.T) {
		a := ")]"
		out := ParenthesesIsValid(a)
		assert.False(t, out)
	})

	t.Run("invalid", func(t *testing.T) {
		a := "((((([{()}[]]]{{{[]}}})))))"
		out := ParenthesesIsValid(a)
		assert.False(t, out)
	})
}
