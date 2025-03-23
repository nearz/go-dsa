package sort

import "testing"

func TestCountSort(t *testing.T) {
	s := []int{0, 0, 1, 1, 9, 7, 8, 1, 4, 5, 2}
	CountSort(s)
}
