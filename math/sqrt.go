package math

//Sqrt return floor(sqrt(A))
func Sqrt(a int) int {
	start, end := 1, a
	for start <= end {
		mid := start + (end-start)/2
		div := a / mid
		if div == mid || div == mid-1 {
			return div
		} else if mid*mid < a {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return 0
}
