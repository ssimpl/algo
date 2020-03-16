package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSwapPairs(t *testing.T) {
	t.Run("one element", func(t *testing.T) {
		a := &node{
			value: 1,
			next:  nil,
		}
		exp := &node{
			value: 1,
			next:  nil,
		}
		out := SwapPairs(a)
		assert.Equal(t, exp, out)
	})

	t.Run("even number of elements", func(t *testing.T) {
		a := &node{
			value: 1,
			next: &node{
				value: 2,
				next: &node{
					value: 3,
					next: &node{
						value: 4,
						next:  nil,
					},
				},
			},
		}
		exp := &node{
			value: 2,
			next: &node{
				value: 1,
				next: &node{
					value: 4,
					next: &node{
						value: 3,
						next:  nil,
					},
				},
			},
		}
		out := SwapPairs(a)
		assert.Equal(t, exp, out)
	})

	t.Run("odd number of elements", func(t *testing.T) {
		a := &node{
			value: 1,
			next: &node{
				value: 2,
				next: &node{
					value: 3,
					next: &node{
						value: 4,
						next: &node{
							value: 5,
							next:  nil,
						},
					},
				},
			},
		}
		exp := &node{
			value: 2,
			next: &node{
				value: 1,
				next: &node{
					value: 4,
					next: &node{
						value: 3,
						next: &node{
							value: 5,
							next:  nil,
						},
					},
				},
			},
		}
		out := SwapPairs(a)
		assert.Equal(t, exp, out)
	})
}
