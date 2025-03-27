package sort

import "cmp"

func MergeSort[S ~[]T, T cmp.Ordered](arr S) S {
	if len(arr) < 2 {
		return arr
	}
	l := MergeSort(arr[:len(arr)/2])
	r := MergeSort(arr[len(arr)/2:])
	return merge(l, r)
}

func merge[S ~[]T, T cmp.Ordered](l, r S) S {
	i := 0
	j := 0
	final := []T{}
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

func MergeSortFunc[S ~[]T, T any](arr S, cmp func(a, b T) bool) S {
	if len(arr) < 2 {
		return arr
	}
	l := MergeSortFunc(arr[:len(arr)/2], cmp)
	r := MergeSortFunc(arr[len(arr)/2:], cmp)
	return mergeFunc(l, r, cmp)
}

func mergeFunc[S ~[]T, T any](l, r S, cmp func(a, b T) bool) S {
	i := 0
	j := 0
	final := []T{}
	for i < len(l) && j < len(r) {
		if cmp(l[i], r[j]) {
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
