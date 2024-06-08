package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTopKFrequent(t *testing.T) {
	t.Parallel()

	t.Run("base case 1", func(t *testing.T) {
		t.Parallel()

		res := topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2)
		assert.Equal(t, []int{1, 2}, res)
	})
	t.Run("base case 2", func(t *testing.T) {
		t.Parallel()

		res := topKFrequent([]int{1}, 1)
		assert.Equal(t, []int{1}, res)
	})
}
