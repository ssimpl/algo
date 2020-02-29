package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindRotationCount(t *testing.T) {
	t.Run("zero items", func(t *testing.T) {
		var arr []int
		expected := 0
		output := FindRotationCount(arr, 0, len(arr)-1)
		assert.Equal(t, expected, output)
	})

	t.Run("one item", func(t *testing.T) {
		arr := []int{1}
		expected := 0
		output := FindRotationCount(arr, 0, len(arr)-1)
		assert.Equal(t, expected, output)
	})

	t.Run("not rotated", func(t *testing.T) {
		arr := []int{1, 3, 5, 6, 7, 8, 12, 15, 22}
		expected := 0
		output := FindRotationCount(arr, 0, len(arr)-1)
		assert.Equal(t, expected, output)
	})

	t.Run("rotated", func(t *testing.T) {
		arr := []int{12, 15, 22, 1, 3, 5, 6, 7, 8}
		expected := 3
		output := FindRotationCount(arr, 0, len(arr)-1)
		assert.Equal(t, expected, output)
	})

	t.Run("rotated with duplicate", func(t *testing.T) {
		arr := []int{12, 15, 15, 22, 22, 22, 1, 1, 3, 3, 3, 5, 6, 7, 7, 8}
		expected := 6
		output := FindRotationCount(arr, 0, len(arr)-1)
		assert.Equal(t, expected, output)
	})
}
