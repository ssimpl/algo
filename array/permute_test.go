package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPermute(t *testing.T) {
	t.Run("one item", func(t *testing.T) {
		a := []int{1}
		exp := [][]int{
			{1},
		}
		out := Permute(a)
		assert.ElementsMatch(t, exp, out)
	})

	t.Run("two items", func(t *testing.T) {
		a := []int{1, 2}
		exp := [][]int{
			{1, 2},
			{2, 1},
		}
		out := Permute(a)
		assert.ElementsMatch(t, exp, out)
	})

	t.Run("three items", func(t *testing.T) {
		a := []int{1, 2, 3}
		exp := [][]int{
			{1, 2, 3},
			{1, 3, 2},
			{2, 1, 3},
			{2, 3, 1},
			{3, 1, 2},
			{3, 2, 1},
		}
		out := Permute(a)
		assert.ElementsMatch(t, exp, out)
	})

	t.Run("four items", func(t *testing.T) {
		a := []int{1, 2, 3, 4}
		exp := [][]int{
			{1, 2, 3, 4},
			{1, 2, 4, 3},
			{1, 3, 2, 4},
			{1, 3, 4, 2},
			{1, 4, 2, 3},
			{1, 4, 3, 2},
			{2, 1, 3, 4},
			{2, 1, 4, 3},
			{2, 3, 1, 4},
			{2, 3, 4, 1},
			{2, 4, 1, 3},
			{2, 4, 3, 1},
			{3, 1, 2, 4},
			{3, 1, 4, 2},
			{3, 2, 1, 4},
			{3, 2, 4, 1},
			{3, 4, 1, 2},
			{3, 4, 2, 1},
			{4, 1, 2, 3},
			{4, 1, 3, 2},
			{4, 2, 1, 3},
			{4, 2, 3, 1},
			{4, 3, 1, 2},
			{4, 3, 2, 1},
		}
		out := Permute(a)
		assert.ElementsMatch(t, exp, out)
	})
}
