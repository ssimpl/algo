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
