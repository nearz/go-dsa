package sort

import (
	"fmt"
	"slices"
	"testing"

	arrayutils "github.com/nearz/go-dsa/arrays/arrayUtils"
)

func TestQuickSortInt(t *testing.T) {
	ts := []int{0, 1, 2, 3, 4, 6, 7, 8}
	s := []int{8, 2, 6, 4, 0, 1, 7, 3}
	QuickSort(s)
	if !slices.Equal(ts, s) {
		t.Error("Slice not sorted, Test 1")
	}

	for range 10 {
		r, s := arrayutils.RandAndSorted(100, 1000)
		QuickSort(r)
		if !slices.Equal(r, s) {
			t.Error("Slice not sorted correctly\n")
			fmt.Printf("Sorted Slice: %v\n", s)
			fmt.Printf("Unsorted Slice: %v\n", r)
		}
	}
}

func TestQuickSortStr(t *testing.T) {
	ts := []string{"apple", "cat", "food", "moon", "test", "zebra"}
	s := []string{"moon", "apple", "test", "cat", "zebra", "food"}
	QuickSort(s)
	if !slices.Equal(ts, s) {
		t.Error("Slice not sorted, Test 1 Strings")
	}
}

func TestQuickSortFlt(t *testing.T) {
	ts := []float64{-2.4, -1.9, -1.2, 0.1, 0.7, 1.1, 1.7, 2.3, 3.1}
	s := []float64{2.3, 1.1, -1.2, 0.1, -1.9, 3.1, -2.4, 0.7, 1.7}
	QuickSort(s)
	if !slices.Equal(ts, s) {
		t.Error("Slice not sorted, Test 1 Floats")
	}
}

func TestQuickSortCustomType(t *testing.T) {
	type CInts []int
	ci := CInts{8, 2, 6, 4, 0, 1, 7, 3}
	ts := CInts{0, 1, 2, 3, 4, 6, 7, 8}
	QuickSort(ci)
	if !slices.Equal(ts, ci) {
		t.Error("Slice not sorted, Test 1 Custom Type")
	}
}

func TestQuickSortFunc(t *testing.T) {
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
	QuickSortFunc(ps, func(a, b P) bool {
		return a.age <= b.age
	})
	if !slices.Equal(tps, ps) {
		t.Error("Slice not sorted, Test 1 structs")
	}
}

func TestQuickSortFuncCustomType(t *testing.T) {
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
	QuickSortFunc(ps, func(a, b P) bool {
		return a.age <= b.age
	})
	if !slices.Equal(tps, ps) {
		t.Error("Slice not sorted, Test 1 custom slice structs")
	}
}

func TestQuickSortFuncStr(t *testing.T) {
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
	QuickSortFunc(ps, func(a, b P) bool {
		return a.name < b.name
	})
	if !slices.Equal(tps, ps) {
		t.Error("Slice not sorted, Test 1 structs compare string")
	}
}
