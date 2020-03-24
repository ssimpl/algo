package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNextGreater(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		a := []int{4, 5, 2, 10}
		exp := []int{5, 10, 10, -1}
		out := NextGreater(a)
		assert.Equal(t, exp, out)
	})

	t.Run("case 2", func(t *testing.T) {
		a := []int{3, 2, 1}
		exp := []int{-1, -1, -1}
		out := NextGreater(a)
		assert.Equal(t, exp, out)
	})

	t.Run("case 3", func(t *testing.T) {
		a := []int{34, 35, 27, 42, 5, 28, 39, 20, 28}
		exp := []int{35, 42, 42, -1, 28, 39, -1, 28, -1}
		out := NextGreater(a)
		assert.Equal(t, exp, out)
	})
}
