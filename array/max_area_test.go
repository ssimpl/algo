package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxArea(t *testing.T) {
	a := []int{1, 5, 4, 3}
	exp := 6
	out := MaxArea(a)
	assert.Equal(t, exp, out)
}
