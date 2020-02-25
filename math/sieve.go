package math

import "math"

func Sieve(n int) []int {
	prime := make([]int, n+1)
	for i := 2; i <= n; i++ {
		prime[i] = 1
	}

	sqrtN := int(math.Sqrt(float64(n)))
	for i := 2; i <= sqrtN; i++ {
		if prime[i] == 1 {
			for j := 2; i*j <= n; j++ {
				prime[i*j] = 0
			}
		}
	}

	var res []int
	for i := 2; i <= n; i++ {
		if prime[i] == 1 {
			res = append(res, i)
		}
	}

	return res
}
