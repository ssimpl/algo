package array

import (
	"math"
)

func MaxArea(a []int) int {
	var i, j, max int

	j = len(a) - 1
	for i != j {
		cur := (j - i) * int(math.Min(float64(a[i]), float64(a[j])))
		if cur > max {
			max = cur
		}
		if a[i] < a[j] {
			i++
		} else {
			j--
		}
	}

	return max
}
