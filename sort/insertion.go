package sort

func InsertionSort(a []int) []int {
	var curVal, prevPos int
	for i := 1; i < len(a); i++ {
		curVal = a[i]
		prevPos = i - 1
		for prevPos >= 0 && a[prevPos] > curVal {
			a[prevPos+1] = a[prevPos]
			prevPos--
		}
		if prevPos != i-1 {
			a[prevPos+1] = curVal
		}
	}
	return a
}
