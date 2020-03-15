package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntersect(t *testing.T) {
	t.Run("one is empty", func(t *testing.T) {
		a := []int{1, 2, 3, 3, 4, 5, 6}
		var b []int
		var exp []int
		out := Intersect(a, b)
		assert.Equal(t, exp, out)
	})

	t.Run("duplicates", func(t *testing.T) {
		a := []int{1, 2, 3, 3, 4, 5, 6}
		b := []int{0, 3, 5}
		exp := []int{3, 5}
		out := Intersect(a, b)
		assert.Equal(t, exp, out)
	})

	t.Run("same number of duplicates", func(t *testing.T) {
		a := []int{1, 2, 3, 3, 4, 5, 6}
		b := []int{-2, -1, 3, 3, 5}
		exp := []int{3, 3, 5}
		out := Intersect(a, b)
		assert.Equal(t, exp, out)
	})
}
