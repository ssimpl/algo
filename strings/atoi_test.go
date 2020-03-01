package strings

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAtoi(t *testing.T) {
	t.Run("with internal space", func(t *testing.T) {
		s := "1 2345"
		expected := 1
		output := Atoi(s)
		assert.Equal(t, expected, output)
	})

	t.Run("with pre space", func(t *testing.T) {
		s := "  12345"
		expected := 12345
		output := Atoi(s)
		assert.Equal(t, expected, output)
	})

	t.Run("with pre zero", func(t *testing.T) {
		s := "0012345"
		expected := 12345
		output := Atoi(s)
		assert.Equal(t, expected, output)
	})

	t.Run("garbage after the number", func(t *testing.T) {
		s := "1asd"
		expected := 1
		output := Atoi(s)
		assert.Equal(t, expected, output)
	})

	t.Run("no numeric character before garbage", func(t *testing.T) {
		s := "a2"
		expected := 0
		output := Atoi(s)
		assert.Equal(t, expected, output)
	})

	t.Run("positive overflow", func(t *testing.T) {
		s := "11111111111"
		expected := math.MaxInt32
		output := Atoi(s)
		assert.Equal(t, expected, output)
	})

	t.Run("negative number", func(t *testing.T) {
		s := "-111"
		expected := -111
		output := Atoi(s)
		assert.Equal(t, expected, output)
	})

	t.Run("negative number with pre space", func(t *testing.T) {
		s := "  -111"
		expected := -111
		output := Atoi(s)
		assert.Equal(t, expected, output)
	})

	t.Run("negative overflow", func(t *testing.T) {
		s := "-11111111111"
		expected := math.MinInt32
		output := Atoi(s)
		assert.Equal(t, expected, output)
	})

	t.Run("positive number", func(t *testing.T) {
		s := "+111"
		expected := 111
		output := Atoi(s)
		assert.Equal(t, expected, output)
	})
}
