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

func PascalTriangleRow(n int) []int {
	row := make([]int, n+1)
	row[0], row[n] = 1, 1

	for i := 0; i < n/2; i++ {
		x := row[i] * (n - i) / (i + 1)
		row[i+1], row[n-1-i] = x, x
	}

	return row
}
