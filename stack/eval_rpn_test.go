package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvalRPN(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		var a []string
		exp := 0
		out := EvalRPN(a)
		assert.Equal(t, exp, out)
	})

	t.Run("one operand", func(t *testing.T) {
		a := []string{"2"}
		exp := 2
		out := EvalRPN(a)
		assert.Equal(t, exp, out)
	})

	t.Run("several operands", func(t *testing.T) {
		a := []string{"5", "1", "2", "+", "4", "*", "+", "3", "-"}
		exp := 14
		out := EvalRPN(a)
		assert.Equal(t, exp, out)
	})
}
