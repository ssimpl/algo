package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindCount(t *testing.T) {
	arr := []int{1, 1, 3, 3, 5, 5, 5, 5, 5, 9, 9, 11}
	expected := 5
	output := FindCount(arr, 5)
	assert.Equal(t, expected, output)
}
