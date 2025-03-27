package sort

import (
	"fmt"
	"slices"
	"testing"

	arrayutils "github.com/nearz/go-dsa/arrays/arrayUtils"
)

func TestMergeSort(t *testing.T) {
	ts := []int{0, 1, 2, 3, 4, 6, 7, 8}
	s := []int{8, 2, 6, 4, 0, 1, 7, 3}
	f := MergeSort(s)
	if !slices.Equal(ts, f) {
		t.Error("Slice not sorted, Test 1")
	}

	for range 10 {
		r, s := arrayutils.RandAndSorted(100, 1000)
		f := MergeSort(r)
		if !slices.Equal(s, f) {
			t.Error("Slice not sorted correctly\n")
			fmt.Printf("Sorted Slice: %v\n", s)
			fmt.Printf("Unsorted Slice: %v\n", r)
		}
	}
}

func TestMergeSortStr(t *testing.T) {
	ts := []string{"apple", "cat", "food", "moon", "test", "zebra"}
	s := []string{"moon", "apple", "test", "cat", "zebra", "food"}
	r := MergeSort(s)
	if !slices.Equal(ts, r) {
		t.Error("Slice not sorted, Test 1 Strings")
	}
}

func TestMergeSortFlt(t *testing.T) {
	ts := []float64{-2.4, -1.9, -1.2, 0.1, 0.7, 1.1, 1.7, 2.3, 3.1}
	s := []float64{2.3, 1.1, -1.2, 0.1, -1.9, 3.1, -2.4, 0.7, 1.7}
	r := MergeSort(s)
	if !slices.Equal(ts, r) {
		t.Error("Slice not sorted, Test 1 Floats")
	}
}

func TestMergeSortCustomType(t *testing.T) {
	type CInts []int
	ci := CInts{8, 2, 6, 4, 0, 1, 7, 3}
	ts := CInts{0, 1, 2, 3, 4, 6, 7, 8}
	r := MergeSort(ci)
	if !slices.Equal(ts, r) {
		t.Error("Slice not sorted, Test 1 Custom Type")
	}
}

func TestMergeSortFunc(t *testing.T) {
	type P struct {
		name string
		age  int
	}
	ps := []P{
		{"Pat", 36},
		{"Jay", 18},
		{"April", 54},
		{"May", 15},
	}
	tps := []P{
		{"May", 15},
		{"Jay", 18},
		{"Pat", 36},
		{"April", 54},
	}
	r := MergeSortFunc(ps, func(a, b P) bool {
		return a.age < b.age
	})
	if !slices.Equal(tps, r) {
		t.Error("Slice not sorted, Test 1 structs")
	}
}

func TestMergeSortFuncCustomType(t *testing.T) {
	type P struct {
		name string
		age  int
	}
	type Pslice []P
	ps := Pslice{
		{"Pat", 36},
		{"Jay", 18},
		{"April", 54},
		{"May", 15},
	}
	tps := Pslice{
		{"May", 15},
		{"Jay", 18},
		{"Pat", 36},
		{"April", 54},
	}
	r := MergeSortFunc(ps, func(a, b P) bool {
		return a.age <= b.age
	})
	if !slices.Equal(tps, r) {
		t.Error("Slice not sorted, Test 1 custom slice structs")
	}
}

func TestMergeSortFuncStr(t *testing.T) {
	type P struct {
		name string
		age  int
	}
	ps := []P{
		{"Pat", 36},
		{"Jay", 18},
		{"April", 54},
		{"May", 15},
	}
	tps := []P{
		{"April", 54},
		{"Jay", 18},
		{"May", 15},
		{"Pat", 36},
	}
	r := MergeSortFunc(ps, func(a, b P) bool {
		return a.name < b.name
	})
	if !slices.Equal(tps, r) {
		t.Error("Slice not sorted, Test 1 structs compare string")
	}
}
