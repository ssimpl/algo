package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPow(t *testing.T) {
	t.Run("zero", func(t *testing.T) {
		a := 2
		n := 0
		expected := 1
		output := Pow(a, n)
		assert.Equal(t, expected, output)
	})

	t.Run("odd", func(t *testing.T) {
		a := 2
		n := 5
		expected := 32
		output := Pow(a, n)
		assert.Equal(t, expected, output)
	})

	t.Run("even", func(t *testing.T) {
		a := 2
		n := 6
		expected := 64
		output := Pow(a, n)
		assert.Equal(t, expected, output)
	})
}
