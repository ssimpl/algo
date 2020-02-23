package sort

func MergeSort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	mid := len(a) / 2
	return merge(MergeSort(a[:mid]), MergeSort(a[mid:]))
}

func merge(a []int, b []int) []int {
	res := make([]int, 0, len(a)+len(b))

	var i, j int
	for i < len(a) || j < len(b) {
		if i < len(a) && j < len(b) {
			if a[i] < b[j] {
				res = append(res, a[i])
				i++
			} else if b[j] < a[i] {
				res = append(res, b[j])
				j++
			} else {
				res = append(res, a[i], b[j])
				i++
				j++
			}
		} else if i < len(a) {
			res = append(res, a[i:]...)
			break
		} else {
			res = append(res, b[j:]...)
			break
		}
	}

	return res
}
