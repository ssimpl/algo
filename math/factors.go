package math

import (
	"math"
	"sort"
)

func AllFactors(n int) []int {
	var res []int
	sqrtN := int(math.Sqrt(float64(n)))
	for i := 1; i <= sqrtN; i++ {
		if n%i == 0 {
			res = append(res, i)
			if i != n/i {
				res = append(res, n/i)
			}
		}
	}
	sort.Ints(res)
	return res
}
