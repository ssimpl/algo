package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckSubarraySum(t *testing.T) {
	t.Parallel()

	t.Run("base case 1", func(t *testing.T) {
		res := checkSubarraySum([]int{23, 2, 4, 6, 7}, 6)
		assert.True(t, res)
	})
	t.Run("base case 2", func(t *testing.T) {
		res := checkSubarraySum([]int{23, 2, 6, 4, 7}, 6)
		assert.True(t, res)
	})
	t.Run("base case 3", func(t *testing.T) {
		res := checkSubarraySum([]int{23, 2, 6, 4, 7}, 13)
		assert.False(t, res)
	})
	t.Run("base case 4", func(t *testing.T) {
		res := checkSubarraySum([]int{1, 0}, 2)
		assert.False(t, res)
	})
}
