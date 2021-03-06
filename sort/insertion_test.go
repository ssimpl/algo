package sort

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestInsertionSort(t *testing.T) {
	for _, s := range tests {
		input := s[0]
		expected := s[1]
		output := InsertionSort(input)
		assert.Equal(t, expected, output)
	}
}

func BenchmarkInsertionSort(b *testing.B) {
	data := make([]int, benchListSize)
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
