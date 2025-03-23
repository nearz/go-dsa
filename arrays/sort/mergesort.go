package sort

func MergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	l := MergeSort(arr[:len(arr)/2])
	r := MergeSort(arr[len(arr)/2:])
	return merge(l, r)
}

func merge(l []int, r []int) []int {
	i := 0
	j := 0
	final := []int{}
	for i < len(l) && j < len(r) {
		if l[i] < r[j] {
			final = append(final, l[i])
			i++
		} else {
			final = append(final, r[j])
			j++
		}
	}
	for ; i < len(l); i++ {
		final = append(final, l[i])
	}
	for ; j < len(r); j++ {
		final = append(final, r[j])
	}
	return final
}
