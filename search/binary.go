package search

func BinarySearch(arr []int, x, start, end int, searchFirst bool) int {
	if start > end {
		return -1
	}

	mid := start + (end-start)/2
	if arr[mid] == x {
		res := -1
		if searchFirst {
			res = BinarySearch(arr, x, start, mid-1, searchFirst)
		} else {
			res = BinarySearch(arr, x, mid+1, end, searchFirst)
		}
		if res == -1 {
			return mid
		}
		return res
	} else if arr[mid] > x {
		return BinarySearch(arr, x, start, mid-1, searchFirst)
	} else {
		return BinarySearch(arr, x, mid+1, end, searchFirst)
	}
}
