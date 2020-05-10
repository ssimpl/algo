package array

func DNums(a []int, k int) []int {
	var res []int
	if k > len(a) {
		return res
	}

	var dCount int
	vals := make(map[int]int, len(a))

	for _, v := range a[:k] {
		if _, ok := vals[v]; !ok {
			dCount++
		}
		vals[v]++
	}

	res = append(res, dCount)

	for i := k; i < len(a); i++ {
		if v := vals[a[i-k]]; v == 1 {
			dCount--
		}
		vals[a[i-k]]--

		if v := vals[a[i]]; v == 0 {
			dCount++
		}
		vals[a[i]]++

		res = append(res, dCount)
	}

	return res
}
