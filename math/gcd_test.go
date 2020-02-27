package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGCD(t *testing.T) {
	a := 6
	b := 9
	output := GCD(a, b)
	expected := 3
	assert.Equal(t, expected, output)
}
