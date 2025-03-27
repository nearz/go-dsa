package sort

import (
	"cmp"

	arrayutils "github.com/nearz/go-dsa/arrays/arrayUtils"
)

func QuickSort[S ~[]T, T cmp.Ordered](arr S) {
	arrayutils.Shuffle(arr)
	quickSort(arr, 0, len(arr)-1)
}

func quickSort[S ~[]T, T cmp.Ordered](arr S, lo, hi int) {
	if lo < hi {
		pivIdx := partition(arr, hi, lo)
		quickSort(arr, lo, pivIdx-1)
		quickSort(arr, pivIdx+1, hi)
	}
}

func partition[S ~[]T, T cmp.Ordered](arr S, hi, lo int) int {
	piv := arr[hi]
	i := lo
	for j := lo; j < hi; j++ {
		if arr[j] <= piv {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[hi] = arr[hi], arr[i]
	return i
}

func QuickSortFunc[S ~[]T, T any](arr S, cmp func(a, b T) bool) {
	quickSortFunc(arr, 0, len(arr)-1, cmp)
}

func quickSortFunc[S ~[]T, T any](arr S, lo, hi int, cmp func(a, b T) bool) {
	if lo < hi {
		pivIdx := partitionFunc(arr, hi, lo, cmp)
		quickSortFunc(arr, lo, pivIdx-1, cmp)
		quickSortFunc(arr, pivIdx+1, hi, cmp)
	}
}

func partitionFunc[S ~[]T, T any](arr S, hi, lo int, cmp func(a, b T) bool) int {
	piv := arr[hi]
	i := lo
	for j := lo; j < hi; j++ {
		if cmp(arr[j], piv) {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[hi] = arr[hi], arr[i]
	return i
}
