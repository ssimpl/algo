package array

func zip(s1, s2 []int) [][2]int {
	minLen := len(s1)
	if len(s2) < minLen {
		minLen = len(s2)
	}

	res := make([][2]int, 0, minLen)
	for i := 0; i < minLen; i++ {
		res = append(res, [2]int{s1[i], s2[i]})
	}

	return res
}
