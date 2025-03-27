package arrayutils

import (
	"cmp"
	"math/rand"
	"sort"
	"time"
)

func Shuffle[S ~[]T, T cmp.Ordered](arr S) {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	n := len(arr)
	for i := n - 1; i > 0; i-- {
		j := r.Intn(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func RandAndSorted(n, max int) ([]int, []int) {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	randInts := make([]int, n)
	for i := 0; i < n; i++ {
		randInts[i] = r.Intn(2*max) - max
	}
	sortInts := make([]int, n)
	copy(sortInts, randInts)
	sort.Ints(sortInts)
	return randInts, sortInts
}

func MaxInt(arr []int) int {
	var m int
	for i, v := range arr {
		if i == 0 || v > m {
			m = v
		}
	}
	return m
}

func MinInt(arr []int) int {
	var m int
	for i, v := range arr {
		if i == 0 || v < m {
			m = v
		}
	}
	return m
}
