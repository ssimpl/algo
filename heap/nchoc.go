package heap

import (
	"container/heap"
)

const nChocMod = 1000000007

func NChoc(k int, a []int) int {
	var res int

	h := &MaxIntHeap{a}
	heap.Init(h)

	for i := 0; i < k; i++ {
		v := heap.Pop(h).(int)
		res += v
		div := v / 2
		heap.Push(h, div)
	}

	return res % nChocMod
}
