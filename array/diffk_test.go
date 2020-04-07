package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiffPossible(t *testing.T) {
	a := []int{1, 3, 5}
	k := 4
	exp := 1
	out := DiffPossible(a, k)
	assert.Equal(t, exp, out)
}

func TestDiffPossibleWithMap(t *testing.T) {
	t.Run("one item", func(t *testing.T) {
		a := []int{0}
		k := 0
		exp := 0
		out := DiffPossibleWithMap(a, k)
		assert.Equal(t, exp, out)
	})

	t.Run("three different items", func(t *testing.T) {
		a := []int{1, 5, 3}
		k := 2
		exp := 1
		out := DiffPossibleWithMap(a, k)
		assert.Equal(t, exp, out)
	})
}
