package array

func commonChars(words []string) []string {
	m := make(map[byte][]int)
	for i := 0; i < len(words); i++ {
		w := words[i]
		for j := 0; j < len(w); j++ {
			c := w[j]
			if len(m[c]) == 0 {
				m[c] = make([]int, len(words))
			}
			m[c][i]++
		}
	}

	var res []string
	for c, ww := range m {
		min := ww[0]
		for _, n := range ww {
			if n < min {
				min = n
			}
			if n == 0 {
				break
			}
		}

		for i := 0; i < min; i++ {
			res = append(res, string(c))
		}
	}

	return res
}
