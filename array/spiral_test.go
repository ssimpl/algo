package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpiralOrder(t *testing.T) {
	input := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
	expected := []int{1, 2, 3, 4, 8, 12, 16, 15, 14, 13, 9, 5, 6, 7, 11, 10}
	output := SpiralOrder(input)
	assert.Equal(t, expected, output)
}
