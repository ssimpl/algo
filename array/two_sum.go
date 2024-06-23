package array

import "sort"

// 1. Two Sum
// Отличие от исходной задачи в том, что возвращаются числа, а не индексы
func twoSum(nums []int, target int) []int {
	sort.Ints(nums)

	l, r := 0, len(nums)-1
	for l < r {
		n := nums[l] + nums[r]
		if n == target {
			return []int{nums[l], nums[r]}
		} else if n > target {
			r--
		} else {
			l++
		}
	}

	return []int{}
}
