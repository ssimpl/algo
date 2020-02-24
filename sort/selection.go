package sort

func SelectionSort(a []int) []int {
	var minIdx int
	for i := 0; i < len(a)-1; i++ {
		minIdx = i
		for j := i + 1; j < len(a); j++ {
			if a[j] < a[minIdx] {
				minIdx = j
			}
		}
		a[i], a[minIdx] = a[minIdx], a[i]
	}
	return a
}
