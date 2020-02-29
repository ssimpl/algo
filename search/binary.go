package search

func BinarySearch(arr []int, x, start, end int) int {
	if start > end {
		return -1
	}

	mid := start + (end-start)/2
	if arr[mid] == x {
		return mid
	} else if arr[mid] > x {
		return BinarySearch(arr, x, start, mid-1)
	} else {
		return BinarySearch(arr, x, mid+1, end)
	}
}
