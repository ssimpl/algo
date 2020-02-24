package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAllFactors(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		input := 1
		expected := []int{1}
		output := AllFactors(input)
		assert.Equal(t, expected, output)
	})

	t.Run("38808", func(t *testing.T) {
		input := 38808
		expected := []int{1, 2, 3, 4, 6, 7, 8, 9, 11, 12, 14, 18, 21, 22, 24, 28, 33, 36, 42, 44, 49, 56, 63, 66, 72, 77, 84, 88, 98, 99, 126, 132, 147, 154, 168, 196, 198, 231, 252, 264, 294, 308, 392, 396, 441, 462, 504, 539, 588, 616, 693, 792, 882, 924, 1078, 1176, 1386, 1617, 1764, 1848, 2156, 2772, 3234, 3528, 4312, 4851, 5544, 6468, 9702, 12936, 19404, 38808}
		output := AllFactors(input)
		assert.Equal(t, expected, output)
	})
}
