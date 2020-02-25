package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPrime(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		input := 1
		output := IsPrime(input)
		assert.Equal(t, false, output)
	})

	t.Run("7", func(t *testing.T) {
		input := 7
		output := IsPrime(input)
		assert.Equal(t, true, output)
	})

	t.Run("39601", func(t *testing.T) {
		input := 39601
		output := IsPrime(input)
		assert.Equal(t, false, output)
	})

	t.Run("41443", func(t *testing.T) {
		input := 41443
		output := IsPrime(input)
		assert.Equal(t, true, output)
	})
}
