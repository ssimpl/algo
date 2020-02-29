package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSqrt(t *testing.T) {
	t.Run("one", func(t *testing.T) {
		input := 1
		expected := 1
		output := Sqrt(input)
		assert.Equal(t, expected, output)
	})

	t.Run("two", func(t *testing.T) {
		input := 2
		expected := 1
		output := Sqrt(input)
		assert.Equal(t, expected, output)
	})

	t.Run("nine", func(t *testing.T) {
		input := 9
		expected := 3
		output := Sqrt(input)
		assert.Equal(t, expected, output)
	})

	t.Run("eleven", func(t *testing.T) {
		input := 11
		expected := 3
		output := Sqrt(input)
		assert.Equal(t, expected, output)
	})
}
