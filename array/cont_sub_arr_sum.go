package array

// 523. Continuous Subarray Sum
func checkSubarraySum(nums []int, k int) bool {
	m := map[int]int{0: -1}

	var prefSum int
	for i, n := range nums {
		prefSum += n
		r := prefSum % k
		if _, ok := m[r]; !ok {
			m[r] = i
		} else if i-m[r] > 1 {
			return true
		}
	}

	return false
}
