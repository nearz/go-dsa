package sort

import (
	"slices"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	ts := []int{0, 1, 2, 3, 4, 6, 7, 8}
	s := []int{8, 2, 6, 4, 0, 1, 7, 3}
	SelectionSort(s)
	if !slices.Equal(ts, s) {
		t.Error("Slice not sorted, Test 1")
	}

	ts = []int{-1, 8, 14, 19, 24, 25, 31, 49, 50, 51, 57, 65}
	s = []int{49, -1, 57, 14, 24, 8, 31, 65, 51, 25, 19, 50}
	SelectionSort(s)
	if !slices.Equal(ts, s) {
		t.Error("Slice not sorted correctly, Test 2")
	}
}
