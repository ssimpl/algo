package sort

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestQuickSort(t *testing.T) {
	for _, s := range tests {
		input := s[0]
		expected := s[1]
		QuickSort(&input, 0, len(input)-1)
		assert.Equal(t, expected, input)
	}
}

func BenchmarkQuickSort(b *testing.B) {
	data := make([]int, benchListSize)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		rand.Seed(time.Now().UnixNano())
		for j := range data {
			data[j] = rand.Int()
		}
		b.StartTimer()
		QuickSort(&data, 0, len(data)-1)
	}
}
