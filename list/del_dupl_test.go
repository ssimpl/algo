package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteDuplicates(t *testing.T) {
	a := &node{
		value: 1,
		next: &node{
			value: 1,
			next: &node{
				value: 2,
				next: &node{
					value: 3,
					next: &node{
						value: 3,
						next:  nil,
					},
				},
			},
		},
	}
	exp := &node{
		value: 1,
		next: &node{
			value: 2,
			next: &node{
				value: 3,
				next:  nil,
			},
		},
	}
	out := DeleteDuplicates(a)
	assert.Equal(t, exp, out)
}
