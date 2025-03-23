package sort

import (
	"fmt"
	"slices"
	"testing"

	arrayutils "github.com/nearz/go-dsa/arrays/arrayUtils"
)

func TestInsertionSort(t *testing.T) {
	ts := []int{0, 1, 2, 3, 4, 6, 7, 8}
	s := []int{8, 2, 6, 4, 0, 1, 7, 3}
	InsertionSort(s)
	if !slices.Equal(ts, s) {
		t.Error("Slice not sorted, Test 1")
	}

	for range 10 {
		r, s := arrayutils.RandAndSorted(100, 1000)
		InsertionSort(r)
		if !slices.Equal(r, s) {
			t.Error("Slice not sorted correctly\n")
			fmt.Printf("Sorted Slice: %v\n", s)
			fmt.Printf("Unsorted Slice: %v\n", r)
		}
	}
}
