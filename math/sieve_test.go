package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSieve(t *testing.T) {
	t.Run("7", func(t *testing.T) {
		input := 7
		expected := []int{2, 3, 5, 7}
		output := Sieve(input)
		assert.Equal(t, expected, output)
	})

	t.Run("15", func(t *testing.T) {
		input := 15
		expected := []int{2, 3, 5, 7, 11, 13}
		output := Sieve(input)
		assert.Equal(t, expected, output)
	})
}
