package search

func FindCount(arr []int, x int) int {
	first := BinarySearch(arr, x, 0, len(arr)-1, true)
	if first == -1 {
		return 0
	}
	last := BinarySearch(arr, x, 0, len(arr)-1, false)
	return last - first + 1
}
