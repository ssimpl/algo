package search

import "sort"

// 347. Top K Frequent Elements
func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}

	sl := make([][2]int, 0, len(m))
	for k, v := range m {
		sl = append(sl, [2]int{k, v})
	}

	sort.Slice(sl, func(i, j int) bool {
		return sl[i][1] > sl[j][1]
	})

	res := make([]int, 0, k)
	for i := 0; i < k; i++ {
		res = append(res, sl[i][0])
	}

	return res
}
