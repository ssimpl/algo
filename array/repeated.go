package array

func RepeatedNumber(a []int) int {
	cnt := make(map[int]int)
	for i := 0; i < len(a); i++ {
		cnt[a[i]]++
	}

	for n, c := range cnt {
		if c > 1 {
			return n
		}
	}

	return -1
}
