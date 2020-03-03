package bit

import "sort"

func FindMinXor(a []int) int {
	sort.Ints(a)
	min := a[0] ^ a[1]
	for i := 1; i < len(a); i++ {
		iterMin := a[i] ^ a[i-1]
		if iterMin < min {
			min = iterMin
		}
	}
	return min
}
