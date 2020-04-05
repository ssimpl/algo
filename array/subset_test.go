package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubsetsWithDup(t *testing.T) {
	t.Run("ordered items with dup", func(t *testing.T) {
		a := []int{1, 2, 2, 3, 4}
		exp := [][]int{
			{},
			{1},
			{1, 2},
			{1, 2, 2},
			{1, 2, 2, 3},
			{1, 2, 2, 3, 4},
			{1, 2, 2, 4},
			{1, 2, 3},
			{1, 2, 3, 4},
			{1, 2, 4},
			{1, 3},
			{1, 3, 4},
			{1, 4},
			{2},
			{2, 2},
			{2, 2, 3},
			{2, 2, 3, 4},
			{2, 2, 4},
			{2, 3},
			{2, 3, 4},
			{2, 4},
			{3},
			{3, 4},
			{4},
		}
		out := SubsetsWithDup(a)
		assert.Equal(t, exp, out)
	})

	t.Run("unordered", func(t *testing.T) {
		a := []int{6, 6, 3, 3, 6, 5}
		exp := [][]int{
			{},
			{3},
			{3, 3},
			{3, 3, 5},
			{3, 3, 5, 6},
			{3, 3, 5, 6, 6},
			{3, 3, 5, 6, 6, 6},
			{3, 3, 6},
			{3, 3, 6, 6},
			{3, 3, 6, 6, 6},
			{3, 5},
			{3, 5, 6},
			{3, 5, 6, 6},
			{3, 5, 6, 6, 6},
			{3, 6},
			{3, 6, 6},
			{3, 6, 6, 6},
			{5},
			{5, 6},
			{5, 6, 6},
			{5, 6, 6, 6},
			{6},
			{6, 6},
			{6, 6, 6},
		}
		out := SubsetsWithDup(a)
		assert.Equal(t, exp, out)
	})
}
