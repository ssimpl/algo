package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseString(t *testing.T) {
	t.Run("empty string", func(t *testing.T) {
		s := ""
		exp := ""
		out := ReverseString(s)
		assert.Equal(t, exp, out)
	})

	t.Run("not empty string", func(t *testing.T) {
		s := "foobarbaz"
		exp := "zabraboof"
		out := ReverseString(s)
		assert.Equal(t, exp, out)
	})
}
