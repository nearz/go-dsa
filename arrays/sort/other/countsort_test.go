package sort

import (
	"slices"
	"testing"
)

func TestCountSort(t *testing.T) {
	ts := []int{0, 0, 1, 1, 1, 2, 4, 5, 7, 8, 9}
	s := []int{0, 7, 1, 5, 9, 0, 8, 1, 4, 1, 2}
	s = CountSort(s)
	if !slices.Equal(ts, s) {
		t.Error("Slice is not sorted. Test 1")
	}
}
