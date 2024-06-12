package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZip(t *testing.T) {
	t.Parallel()

	t.Run("base case", func(t *testing.T) {
		t.Parallel()

		s1, s2 := []int{1, 2, 3}, []int{4, 5, 6, 7, 8}
		res := zip(s1, s2)
		assert.Equal(t, [][2]int{
			{1, 4}, {2, 5}, {3, 6},
		}, res)
	})
}
