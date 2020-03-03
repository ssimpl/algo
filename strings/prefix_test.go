package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestCommonPrefix(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		var a []string
		exp := ""
		out := LongestCommonPrefix(a)
		assert.Equal(t, exp, out)
	})

	t.Run("one item", func(t *testing.T) {
		a := []string{"abc"}
		exp := "abc"
		out := LongestCommonPrefix(a)
		assert.Equal(t, exp, out)
	})

	t.Run("with empty item", func(t *testing.T) {
		a := []string{"abc", "ab", ""}
		exp := ""
		out := LongestCommonPrefix(a)
		assert.Equal(t, exp, out)
	})

	t.Run("regular case", func(t *testing.T) {
		a := []string{"abcd", "abze", "absadqwe"}
		exp := "ab"
		out := LongestCommonPrefix(a)
		assert.Equal(t, exp, out)
	})
}
