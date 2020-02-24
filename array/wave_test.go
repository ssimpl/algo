package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWave(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	expected := []int{2, 1, 4, 3, 5}
	output := Wave(input)
	assert.Equal(t, expected, output)
}
