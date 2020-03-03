package strings

func LongestCommonPrefix(a []string) string {
	if len(a) == 0 {
		return ""
	}
	if len(a) == 1 {
		return a[0]
	}

	short := a[0]
	min := len(short)
	for i := 1; i < len(a); i++ {
		if len(a[i]) < min {
			short = a[i]
			min = len(a[i])
		}
	}
	if min == 0 {
		return ""
	}

	var pref string
	for i := 0; i < min; i++ {
		for j := 1; j < len(a); j++ {
			if a[j][i] != a[0][i] {
				return pref
			}
		}
		pref += string(a[0][i])
	}

	return pref
}
