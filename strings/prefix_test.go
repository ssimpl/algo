package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestCommonPrefix(t *testing.T) {
	t.Parallel()

	t.Run("empty", func(t *testing.T) {
		t.Parallel()

		var a []string
		exp := ""
		out := LongestCommonPrefix(a)
		assert.Equal(t, exp, out)
	})
	t.Run("one item", func(t *testing.T) {
		t.Parallel()

		a := []string{"abc"}
		exp := "abc"
		out := LongestCommonPrefix(a)
		assert.Equal(t, exp, out)
	})
	t.Run("with empty item", func(t *testing.T) {
		t.Parallel()

		a := []string{"abc", "ab", ""}
		exp := ""
		out := LongestCommonPrefix(a)
		assert.Equal(t, exp, out)
	})
	t.Run("regular case", func(t *testing.T) {
		t.Parallel()

		a := []string{"abcd", "abze", "absadqwe"}
		exp := "ab"
		out := LongestCommonPrefix(a)
		assert.Equal(t, exp, out)
	})
	t.Run("no common prefix", func(t *testing.T) {
		t.Parallel()

		a := []string{"dog", "racecar", "car"}
		exp := ""
		out := LongestCommonPrefix(a)
		assert.Equal(t, exp, out)
	})
}

func TestIsPrefixOfWord(t *testing.T) {
	t.Parallel()

	t.Run("base case", func(t *testing.T) {
		t.Parallel()

		res := IsPrefixOfWord("i love eating burger", "burg")
		assert.Equal(t, 4, res)
	})
}
