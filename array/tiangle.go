package array

func PascalTriangle(a int) [][]int {
	res := make([][]int, a)
	for i := 0; i < a; i++ {
		elemCnt := i + 1
		res[i] = make([]int, elemCnt)
		for j := 0; j < elemCnt; j++ {
			if i == 0 {
				res[i][j] = 1
				continue
			}

			prevRow := res[i-1]

			var sum int
			if j == 0 {
				sum += prevRow[j]
			} else {
				sum += prevRow[j-1]
				if j < len(prevRow) {
					sum += prevRow[j]
				}
			}
			res[i][j] = sum
		}
	}
	return res
}
