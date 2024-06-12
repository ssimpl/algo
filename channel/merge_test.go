package channel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	t.Parallel()

	t.Run("base case", func(t *testing.T) {
		ch1 := make(chan int, 3)
		ch1 <- 1
		ch1 <- 2
		ch1 <- 3
		close(ch1)

		ch2 := make(chan int, 2)
		ch2 <- 4
		ch2 <- 5
		close(ch2)

		ch3 := make(chan int, 4)
		ch3 <- 6
		ch3 <- 7
		ch3 <- 8
		ch3 <- 9
		close(ch3)

		res := make([]int, 0, 9)
		outCh := Merge(ch1, ch2, ch3)
		for v := range outCh {
			res = append(res, v)
		}
		assert.ElementsMatch(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, res)
	})
}
