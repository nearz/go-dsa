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
