package array

func SpiralOrder(a [][]int) []int {
	res := make([]int, 0)

	m := len(a)
	n := len(a[0])

	t := 0
	b := m - 1
	r := n - 1
	l := 0

	dir := 0
	for t <= b && l <= r {
		if dir == 0 {
			for i := l; i <= r; i++ {
				res = append(res, a[t][i])
			}
			t++
			dir = 1
		} else if dir == 1 {
			for i := t; i <= b; i++ {
				res = append(res, a[i][r])
			}
			r--
			dir = 2
		} else if dir == 2 {
			for i := r; i >= l; i-- {
				res = append(res, a[b][i])
			}
			b--
			dir = 3
		} else if dir == 3 {
			for i := b; i >= t; i-- {
				res = append(res, a[i][l])
			}
			l++
			dir = 0
		}
	}

	return res
}
