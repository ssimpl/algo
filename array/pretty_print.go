package array

func PrettyPrint(n int) [][]int {
	size := n*2 - 1
	res := make([][]int, size)
	for i := 0; i < n; i++ {
		dif := 0
		res[i] = make([]int, size)
		res[size-i-1] = make([]int, size)
		for j := 0; j < n; j++ {
			val := n - dif
			res[i][j] = val
			res[i][size-j-1] = val
			res[size-i-1][j] = val
			res[size-i-1][size-j-1] = val
			if dif < i {
				dif++
			}
		}
	}
	return res
}
