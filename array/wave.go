package array

import "sort"

func Wave(a []int) []int {
	sort.Ints(a)
	for i := 1; i < len(a); i += 2 {
		a[i-1], a[i] = a[i], a[i-1]
	}
	return a
}
