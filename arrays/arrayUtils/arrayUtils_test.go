package arrayutils

import (
	"fmt"
	"testing"
)

func TestRandAndSorted(t *testing.T) {
	r, s := RandAndSorted(10, 50)
	fmt.Printf("Random Ints: %v\n", r)
	fmt.Printf("Sorted Ints: %v\n", s)
}

func TestMaxInt(t *testing.T) {
	s := []int{4, 2, 1, 9, 6, 0}
	m := MaxInt(s)
	fmt.Printf("Max: %d\n", m)
}

func TestMinInt(t *testing.T) {
	s := []int{4, 2, 1, 9, 6, 0, -4}
	m := MinInt(s)
	fmt.Printf("Max: %d\n", m)
}
