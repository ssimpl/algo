package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnagrams(t *testing.T) {
	a := []string{"cat", "dog", "god", "tca"}
	exp := [][]int{{1, 4}, {2, 3}}
	out := Anagrams(a)
	assert.Equal(t, exp, out)
}
