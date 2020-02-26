package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecimalToBinary(t *testing.T) {
	t.Run("0", func(t *testing.T) {
		input := 0
		expected := "0"
		output := DecimalToBinary(input)
		assert.Equal(t, expected, output)
	})

	t.Run("255", func(t *testing.T) {
		input := 255
		expected := "11111111"
		output := DecimalToBinary(input)
		assert.Equal(t, expected, output)
	})
}
