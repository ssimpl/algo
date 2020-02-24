package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPascalTriangle(t *testing.T) {
	input := 5
	expected := [][]int{
		{1},
		{1, 1},
		{1, 2, 1},
		{1, 3, 3, 1},
		{1, 4, 6, 4, 1},
	}
	output := PascalTriangle(input)
	assert.Equal(t, expected, output)
}

func TestPascalTriangleRow(t *testing.T) {
	input := 4
	expected := []int{1, 4, 6, 4, 1}
	output := PascalTriangleRow(input)
	assert.Equal(t, expected, output)
}
