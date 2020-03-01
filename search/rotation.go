package search

func FindRotationCount(arr []int, start, end int) int {
	if start >= end {
		return start
	}

	mid := start + (end-start)/2
	var count int
	if arr[mid] > arr[end] {
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

func FindInRotated(arr []int, x int) int {
	shift := FindRotationCount(arr, 0, len(arr)-1)
	unshiftedArr := make([]int, 0, len(arr))
	unshiftedArr = append(unshiftedArr, arr[shift:]...)
	unshiftedArr = append(unshiftedArr, arr[:shift]...)
	pos := BinarySearch(unshiftedArr, x, 0, len(unshiftedArr)-1, true)
	if pos == -1 {
		return -1
	}
	return (pos + shift) % len(arr)
}
