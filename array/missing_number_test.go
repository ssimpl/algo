package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMissingNumber(t *testing.T) {
	t.Parallel()

	t.Run("base cases", func(t *testing.T) {
		t.Parallel()

		t.Run("case1", func(t *testing.T) {
			t.Parallel()

			res := missingNumber([]int{3, 0, 1})
			assert.Equal(t, 2, res)
		})
		t.Run("case2", func(t *testing.T) {
			t.Parallel()

			res := missingNumber([]int{0, 1})
			assert.Equal(t, 2, res)
		})
		t.Run("case3", func(t *testing.T) {
			t.Parallel()

			res := missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1})
			assert.Equal(t, 8, res)
		})
	})
}
