package array

// 268. Missing Number
func missingNumber(nums []int) int {
	var sum, target = 0, len(nums)
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		target += i
	}
	return target - sum
}
