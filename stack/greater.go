package stack

func NextGreater(a []int) []int {
	var res []int
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[j] > a[i] {
				res = append(res, a[j])
				break
			} else if j == len(a)-1 {
				res = append(res, -1)
			}
		}
	}

	res = append(res, -1)

	return res
}
