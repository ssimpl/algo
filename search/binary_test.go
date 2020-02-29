package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	t.Run("mid is x", func(t *testing.T) {
		arr := []int{2, 5, 8, 9, 12, 19, 22, 45, 46, 49, 51, 55}
		x := 19
		index := BinarySearch(arr, x, 0, len(arr)-1)
		assert.Equal(t, 5, index)
	})

	t.Run("found", func(t *testing.T) {
		arr := []int{2, 5, 8, 9, 12, 19, 22, 45, 46, 49, 51, 55}
		x := 8
		index := BinarySearch(arr, x, 0, len(arr)-1)
		assert.Equal(t, 2, index)
	})

	t.Run("not found", func(t *testing.T) {
		arr := []int{2, 5, 8, 9, 12, 19, 22, 45, 46, 49, 51, 55}
		x := 13
		index := BinarySearch(arr, x, 0, len(arr)-1)
		assert.Equal(t, -1, index)
	})
}
