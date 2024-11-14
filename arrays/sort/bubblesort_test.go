package array_sort

import (
	"slices"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	a := []int{5, 3, 2, 4, 5, 1}
	b := []int{1, 2, 3, 4, 5, 5}
	BubbleSort(a)
	r := slices.Equal(a, b)
	if !r {
		t.Errorf("BubbleSort = %t, expected true", r)
	}

	a = []int{1, 2}
	b = []int{1, 2}
	BubbleSort(a)
	r = slices.Equal(a, b)
	if !r {
		t.Errorf("BubbleSort = %t, expected true", r)
	}

	a = []int{5, 4, 3, 2, 1, 0}
	b = []int{0, 1, 2, 3, 4, 5}
	BubbleSort(a)
	r = slices.Equal(a, b)
	if !r {
		t.Errorf("BubbleSort = %t, expected true", r)
	}
}
