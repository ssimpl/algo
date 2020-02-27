package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		input := 12121
		output := IsPalindrome(input)
		assert.Equal(t, true, output)
	})

	t.Run("false", func(t *testing.T) {
		input := 123
		output := IsPalindrome(input)
		assert.Equal(t, false, output)
	})
}
