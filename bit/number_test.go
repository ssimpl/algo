package bit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleNumber(t *testing.T) {
	a := []int{1, 2, 2, 3, 1}
	exp := 3
	out := SingleNumber(a)
	assert.Equal(t, exp, out)
}
