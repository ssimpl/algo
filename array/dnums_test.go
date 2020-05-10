package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDNums(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		a := []int{1, 2, 1, 3, 4, 3}
		k := 3
		exp := []int{2, 3, 3, 2}
		out := DNums(a, k)
		assert.Equal(t, exp, out)
	})
}
