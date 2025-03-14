package sort

import (
	"math/rand"
	"time"
)

func QuickSort(arr []int) {
	hi := len(arr) - 1
	lo := 0
	shuffle(arr)
	quickSort(arr, lo, hi)
}

func quickSort(arr []int, lo, hi int) {
	if lo < hi {
		pivIdx := partition(arr, hi, lo)
		quickSort(arr, lo, pivIdx-1)
		quickSort(arr, pivIdx+1, hi)
	}
}

func partition(arr []int, hi, lo int) int {
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

func shuffle(arr []int) {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	n := len(arr)
	for i := n - 1; i > 0; i-- {
		j := r.Intn(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}
}
