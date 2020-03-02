package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseWords(t *testing.T) {
	t.Run("empty string", func(t *testing.T) {
		s := ""
		exp := ""
		out := ReverseWords(s)
		assert.Equal(t, exp, out)
	})

	t.Run("one word", func(t *testing.T) {
		s := "word"
		exp := "word"
		out := ReverseWords(s)
		assert.Equal(t, exp, out)
	})

	t.Run("with after spaces", func(t *testing.T) {
		s := "word   "
		exp := "word"
		out := ReverseWords(s)
		assert.Equal(t, exp, out)
	})

	t.Run("with pre spaces", func(t *testing.T) {
		s := "  the sky is blue"
		exp := "blue is sky the"
		out := ReverseWords(s)
		assert.Equal(t, exp, out)
	})

	t.Run("with double inner spaces", func(t *testing.T) {
		s := "  the  sky   is blue "
		exp := "blue is sky the"
		out := ReverseWords(s)
		assert.Equal(t, exp, out)
	})
}
