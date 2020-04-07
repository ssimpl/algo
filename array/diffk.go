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

func DiffPossibleWithMap(a []int, k int) int {
	values := make(map[int]int, len(a))
	for _, v := range a {
		values[v] += 1
	}
	for _, v := range a {
		exp := v - k
		if count, ok := values[exp]; ok && (exp != v || count > 1) {
			return 1
		}
	}
	return 0
}
