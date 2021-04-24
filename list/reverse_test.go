package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
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
		value: 5,
		next: &node{
			value: 4,
			next: &node{
				value: 3,
				next: &node{
					value: 2,
					next: &node{
						value: 1,
						next:  nil,
					},
				},
			},
		},
	}
	out := Reverse(a)
	assert.Equal(t, exp, out)
}

func TestReverse2(t *testing.T) {
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
		value: 5,
		next: &node{
			value: 4,
			next: &node{
				value: 3,
				next: &node{
					value: 2,
					next: &node{
						value: 1,
						next:  nil,
					},
				},
			},
		},
	}
	out := Reverse2(a, nil)
	assert.Equal(t, exp, out)
}
