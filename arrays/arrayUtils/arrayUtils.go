package arrayutils

import (
	"math/rand"
	"sort"
	"time"
)

func Shuffle(arr []int) {
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
