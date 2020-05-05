package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombine(t *testing.T) {
	t.Run("two items", func(t *testing.T) {
		n := 4
		k := 2
		out := Combine(n, k)
		exp := [][]int{
			{1, 2},
			{1, 3},
			{1, 4},
			{2, 3},
			{2, 4},
			{3, 4},
		}
		assert.Equal(t, exp, out)
	})

	t.Run("three items", func(t *testing.T) {
		n := 4
		k := 3
		out := Combine(n, k)
		exp := [][]int{
			{1, 2, 3},
			{1, 2, 4},
			{1, 3, 4},
			{2, 3, 4},
		}
		assert.Equal(t, exp, out)
	})
}
