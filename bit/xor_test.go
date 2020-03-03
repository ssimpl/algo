package bit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMinXor(t *testing.T) {
	a := []int{12, 4, 6, 2}
	exp := 2
	out := FindMinXor(a)
	assert.Equal(t, exp, out)
}
