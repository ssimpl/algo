package array

func Combine(n int, k int) [][]int {
	var res [][]int
	for i := 1; i <= n; i++ {
		res = append(res, combine(i, n, k)...)
	}
	return res
}

func combine(start int, n int, k int) [][]int {
	if k == 1 {
		return [][]int{{start}}
	}

	var res [][]int
	for i := start; i < n; i++ {
		items := combine(i+1, n, k-1)
		if len(items) == 0 {
			break
		}

		for _, s := range items {
			item := append([]int{start}, s...)
			res = append(res, item)
		}
	}

	return res
}
