package heap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNChoc(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		k := 3
		a := []int{6, 5}
		exp := 14
		out := NChoc(k, a)
		assert.Equal(t, exp, out)
	})
}
