package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestConsecutive(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		a := []int{100, 4, 200, 1, 3, 2}
		exp := 4
		out := LongestConsecutive(a)
		assert.Equal(t, exp, out)
	})

	t.Run("case2", func(t *testing.T) {
		a := []int{6, 4, 5, 2, 3}
		exp := 5
		out := LongestConsecutive(a)
		assert.Equal(t, exp, out)
	})
}
