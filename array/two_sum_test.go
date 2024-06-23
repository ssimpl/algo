package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	t.Parallel()

	t.Run("base cases", func(t *testing.T) {
		t.Parallel()

		t.Run("case1", func(t *testing.T) {
			t.Parallel()

			res := twoSum([]int{2, 7, 11, 1}, 9)
			assert.Equal(t, []int{2, 7}, res)
		})
		t.Run("case2", func(t *testing.T) {
			t.Parallel()

			res := twoSum([]int{3, 2, 4}, 6)
			assert.Equal(t, []int{2, 4}, res)
		})
		t.Run("case3", func(t *testing.T) {
			t.Parallel()

			res := twoSum([]int{3, 3}, 6)
			assert.Equal(t, []int{3, 3}, res)
		})
		t.Run("case4", func(t *testing.T) {
			t.Parallel()

			res := twoSum([]int{1, 2, 3, 4}, 5)
			assert.Equal(t, []int{1, 4}, res)
		})
		t.Run("case5", func(t *testing.T) {
			t.Parallel()

			res := twoSum([]int{4}, 8)
			assert.Equal(t, []int{}, res)
		})
		t.Run("case6", func(t *testing.T) {
			t.Parallel()

			res := twoSum([]int{4, 4}, 8)
			assert.Equal(t, []int{4, 4}, res)
		})
	})
}
