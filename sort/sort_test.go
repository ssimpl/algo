package sort

import (
	"math/rand"
	"sort"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const benchListSize = 10000

var tests = [][][]int{
	{{1, 4, 7, 8, 10}, {1, 4, 7, 8, 10}},
	{{10, 8, 7, 4, 1}, {1, 4, 7, 8, 10}},
	{{25, 17, 31, 13, 2}, {2, 13, 17, 25, 31}},
	{{11, 6, 30, 11, 4, 99}, {4, 6, 11, 11, 30, 99}},
}

func TestBuiltInSort(t *testing.T) {
	for _, s := range tests {
		input := s[0]
		expected := s[1]
		sort.Ints(input)
		assert.Equal(t, expected, input)
	}
}

func BenchmarkBuiltInSort(b *testing.B) {
	data := make([]int, benchListSize)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		rand.Seed(time.Now().UnixNano())
		for j := range data {
			data[j] = rand.Int()
		}
		b.StartTimer()
		sort.Ints(data)
	}
}
