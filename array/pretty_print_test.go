package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrettyPrint(t *testing.T) {
	input := 4
	expected := [][]int{
		{4, 4, 4, 4, 4, 4, 4},
		{4, 3, 3, 3, 3, 3, 4},
		{4, 3, 2, 2, 2, 3, 4},
		{4, 3, 2, 1, 2, 3, 4},
		{4, 3, 2, 2, 2, 3, 4},
		{4, 3, 3, 3, 3, 3, 4},
		{4, 4, 4, 4, 4, 4, 4},
	}
	output := PrettyPrint(input)
	assert.Equal(t, expected, output)
}
