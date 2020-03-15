package array

func DiffPossible(a []int, k int) int {
	for i := len(a) - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if a[i]-a[j] == k {
				return 1
			}
		}
	}
	return 0
}
