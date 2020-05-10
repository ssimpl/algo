package heap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDHeap(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		n := 4
		exp := 3
		out := DHeap(n)
		assert.Equal(t, exp, out)
	})

	t.Run("case 2", func(t *testing.T) {
		n := 10
		exp := 3360
		out := DHeap(n)
		assert.Equal(t, exp, out)
	})

	t.Run("case 3", func(t *testing.T) {
		n := 100
		exp := 812145033
		out := DHeap(n)
		assert.Equal(t, exp, out)
	})
}
