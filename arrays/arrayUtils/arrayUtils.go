package arrayutils

import (
	"math/rand"
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
