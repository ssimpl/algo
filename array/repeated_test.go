package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepeatedNumber(t *testing.T) {
	t.Run("no repeated numbers", func(t *testing.T) {
		a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		exp := -1
		out := RepeatedNumber(a)
		assert.Equal(t, exp, out)
	})

	t.Run("one repeated number", func(t *testing.T) {
		a := []int{3, 4, 1, 4, 5}
		exp := 4
		out := RepeatedNumber(a)
		assert.Equal(t, exp, out)
	})

	t.Run("several repeated numbers", func(t *testing.T) {
		a := []int{3, 4, 1, 4, 5, 1}
		exp := 4
		out := RepeatedNumber(a)
		assert.Equal(t, exp, out)
	})
}
