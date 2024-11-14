package array_search

import "testing"

func TestBinarySearch(t *testing.T) {
	s := []int{2, 6, 7, 10, 15, 20, 32, 53, 77, 101}
	v := 53
	i := BinarySearch(s, v)
	if i != 7 {
		t.Errorf("BinarySearchRecursive = %d; expected: 7", i)
	}

	v = 108
	i = BinarySearch(s, v)
	if i != -1 {
		t.Errorf("BinarySearchRecursive = %d; expected: -1", i)
	}
}
