package arrays

func BinarySearchRecursive(arr []int, v int, lo int, hi int) int {
	m := lo + (hi-lo)/2
	if lo >= hi {
		return -1
	} else if arr[m] == v {
		return m
	} else if v > arr[m] {
		return BinarySearchRecursive(arr, v, m+1, hi)
	} else {
		return BinarySearchRecursive(arr, v, lo, m)
	}
}

func BinarySearch(arr []int, v int) int {
	lo := 0
	hi := len(arr)
	for lo < hi {
		m := lo + (hi-lo)/2
		if arr[m] == v {
			return m
		} else if v > arr[m] {
			lo = m + 1
		} else {
			hi = m
		}
	}
	return -1
}
