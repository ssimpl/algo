package array

func Permute(a []int) [][]int {
	var res = [][]int{a}
	res = append(res, permute(a, 0)...)
	return res
}

func permute(a []int, start int) [][]int {
	var res [][]int
	for i := start; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			newA := make([]int, len(a))
			copy(newA, a)
			newA[i], newA[j] = newA[j], newA[i]
			res = append(res, newA)
			res = append(res, permute(newA, i+1)...)
		}
	}
	return res
}
