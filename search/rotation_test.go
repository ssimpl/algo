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

	t.Run("half the same", func(t *testing.T) {
		arr := []int{1, 1, 3, 3, 3, 5, 6, 7, 7, 7, 7, 7, 7, 7, 7, 7}
		expected := 0
		output := FindRotationCount(arr, 0, len(arr)-1)
		assert.Equal(t, expected, output)
	})

	t.Run("all the same", func(t *testing.T) {
		arr := []int{7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7}
		expected := 0
		output := FindRotationCount(arr, 0, len(arr)-1)
		assert.Equal(t, expected, output)
	})
}

func TestFindInRotated(t *testing.T) {
	arr := []int{180, 181, 182, 183, 184, 187, 188, 189, 191, 192, 193, 194, 195, 196, 201, 202, 203, 204, 3, 4, 5, 6, 7, 8, 9, 10, 14, 16, 17, 18, 19, 23, 26, 27, 28, 29, 32, 33, 36, 37, 38, 39, 41, 42, 43, 45, 48, 51, 52, 53, 54, 56, 62, 63, 64, 67, 69, 72, 73, 75, 77, 78, 79, 83, 85, 87, 90, 91, 92, 93, 96, 98, 99, 101, 102, 104, 105, 106, 107, 108, 109, 111, 113, 115, 116, 118, 119, 120, 122, 123, 124, 126, 127, 129, 130, 135, 137, 138, 139, 143, 144, 145, 147, 149, 152, 155, 156, 160, 162, 163, 164, 166, 168, 169, 170, 171, 172, 173, 174, 175, 176, 177}
	x := 42
	expected := 43
	output := FindInRotated(arr, x)
	assert.Equal(t, expected, output)
}
