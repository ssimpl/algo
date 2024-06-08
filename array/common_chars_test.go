package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommonChars(t *testing.T) {
	t.Parallel()

	t.Run("common case 1", func(t *testing.T) {
		t.Parallel()

		res := commonChars([]string{"bella", "label", "roller"})
		assert.Equal(t, []string{"e", "l", "l"}, res)
	})
	t.Run("common case 2", func(t *testing.T) {
		t.Parallel()

		res := commonChars([]string{"cool", "lock", "cook"})
		assert.Equal(t, []string{"c", "o"}, res)
	})
}
