package search

func FindRotationCount(arr []int, start, end int) int {
	if start >= end {
		return start
	}

	mid := start + (end-start)/2
	var count int
	if arr[mid] >= arr[end] {
		count = FindRotationCount(arr, mid+1, end)
	} else {
		count = FindRotationCount(arr, start, mid-1)
	}

	if arr[count] < arr[mid] {
		return count
	} else if arr[count] > arr[mid] {
		return mid
	}
	if count < mid {
		return count
	}
	return mid
}
