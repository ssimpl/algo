package array

import "math"

func LongestConsecutive(a []int) int {
	values := make(map[int]bool, len(a))
	for _, v := range a {
		values[v] = true
	}

	var max int
	for _, val := range a {
		if !values[val-1] {
			j := val
			for values[j] {
				j++
			}
			max = int(math.Max(float64(max), float64(j-val)))
		}
	}

	return max
}
