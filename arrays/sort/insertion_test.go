package sort

import (
	"slices"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	ts := []int{0, 1, 2, 3, 4, 6, 7, 8}
	s := []int{8, 2, 6, 4, 0, 1, 7, 3}
	InsertionSort(s)
	if !slices.Equal(ts, s) {
		t.Error("Slice not sorted correctly, Test 1")
	}

	ts = []int{5, 8, 14, 19, 24, 25, 31, 49, 50, 51, 57, 65}
	s = []int{49, 5, 57, 14, 24, 8, 31, 65, 51, 25, 19, 50}
	InsertionSort(s)
	if !slices.Equal(ts, s) {
		t.Error("Slice not sorted correctly, Test 2")
	}

	ts = []int{-8, -2, 14, 19, 24, 25, 31, 49, 50, 51, 57, 65}
	s = []int{49, -8, 57, 14, 24, -2, 31, 65, 51, 25, 19, 50}
	InsertionSort(s)
	if !slices.Equal(ts, s) {
		t.Error("Slice not sorted correctly, Test 2")
	}
}
