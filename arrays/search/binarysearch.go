package array_search

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
