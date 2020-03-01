package math

func Pow(a, n int) int {
	if n == 0 {
		return 1
	}
	if n%2 == 1 {
		return Pow(a, n-1) * a
	}
	b := Pow(a, n/2)
	return b * b
}
