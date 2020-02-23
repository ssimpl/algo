package sort

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestInsertionSort(t *testing.T) {
	tests := [][][]int{
		{{1, 4, 7, 8, 10}, {1, 4, 7, 8, 10}},
		{{10, 8, 7, 4, 1}, {1, 4, 7, 8, 10}},
		{{25, 17, 31, 13, 2}, {2, 13, 17, 25, 31}},
		{{11, 6, 30, 11, 4, 99}, {4, 6, 11, 11, 30, 99}},
	}
	for _, s := range tests {
		input := s[0]
		expected := s[1]
		output := InsertionSort(input)
		assert.Equal(t, expected, output)
	}
}

func BenchmarkInsertionSort(b *testing.B) {
	data := make([]int, 10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		rand.Seed(time.Now().UnixNano())
		for j := range data {
			data[j] = rand.Int()
		}
		b.StartTimer()
		InsertionSort(data)
	}
}
