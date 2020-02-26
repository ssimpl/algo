package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzz(t *testing.T) {
	input := 5
	expected := []string{"1", "2", "Fizz", "4", "Buzz"}
	output := FizzBuzz(input)
	assert.Equal(t, expected, output)
}
