package strings

func LengthOfLongestSubstring(a string) int {
	var max, start int

	cc := make(map[string]int, 0)
	for i := 0; i < len(a); i++ {
		c := string(a[i])
		if ci, ok := cc[c]; ok {
			cc[c] = i
			if ci+1 > start {
				start = ci + 1
			}
		} else {
			cc[c] = i
		}

		tmp := len(a[start:i]) + 1
		if tmp > max {
			max = tmp
		}
	}

	return max
}
