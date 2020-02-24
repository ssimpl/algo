package sort

func QuickSort(a []int, start, end int) {
	if start >= end {
		return
	}
	pIndex := partition(a, start, end)
	QuickSort(a, start, pIndex-1)
	QuickSort(a, pIndex+1, end)
}

func partition(a []int, start, end int) int {
	ar := a
	pivot := ar[end]
	pIndex := start
	for i := start; i <= end-1; i++ {
		if ar[i] <= pivot {
			ar[i], ar[pIndex] = ar[pIndex], ar[i]
			pIndex++
		}
	}
	ar[pIndex], ar[end] = ar[end], ar[pIndex]
	return pIndex
}
