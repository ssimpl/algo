package sort

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSelectionSort(t *testing.T) {
	for _, s := range tests {
		input := s[0]
		expected := s[1]
		output := SelectionSort(input)
		assert.Equal(t, expected, output)
	}
}

func BenchmarkSelectionSort(b *testing.B) {
	data := make([]int, benchListSize)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		rand.Seed(time.Now().UnixNano())
		for j := range data {
			data[j] = rand.Int()
		}
		b.StartTimer()
		SelectionSort(data)
	}
}
