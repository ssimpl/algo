package bit

func SingleNumber(a []int) int {
	var res int
	for _, n := range a {
		res ^= n
	}
	return res
}
