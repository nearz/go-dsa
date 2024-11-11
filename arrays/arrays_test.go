package arrays

import (
	"slices"
	"testing"
)

func TestBinarySearchRecursive(t *testing.T) {
	s := []int{2, 6, 7, 10, 15, 20, 32, 53, 77, 101}
	v := 53
	i := BinarySearchRecursive(s, v, 0, len(s))
	if i != 7 {
		t.Errorf("BinarySearchRecursive = %d; expected: 7", i)
	}

	v = 110
	i = BinarySearchRecursive(s, v, 0, len(s))
	if i != -1 {
		t.Errorf("BinarySearchRecursive = %d; expected: -1", i)
	}
}

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

func TestCrystalBall(t *testing.T) {
	b := []bool{false, false, false, false, false, false, false, true, true, true, true, true, true}
	i := CrystalBall(b)
	if i != 7 {
		t.Errorf("CrystalBal = %d; expected 7", i)
	}

	b = []bool{false, false, false, false, false, false, false}
	i = CrystalBall(b)
	if i != -1 {
		t.Errorf("CrystalBall = %d; expected -1", i)
	}

	b = []bool{false, false, false, false, false, false, false, false, false, false, false, true}
	i = CrystalBall(b)
	if i != 11 {
		t.Errorf("CrystalBall = %d, expected 11", i)
	}

	b = []bool{}
	i = CrystalBall(b)
	if i != -1 {
		t.Errorf("CrystalBall = %d, expected -1", i)
	}
}

func TestCrystalBallTwo(t *testing.T) {
	b := []bool{false, false, false, false, false, false, false, true, true, true, true, true, true}
	i := CrystalBallTwo(b)
	if i != 7 {
		t.Errorf("CrystalBal = %d; expected 7", i)
	}

	b = []bool{false, false, false, false, false, false, false}
	i = CrystalBallTwo(b)
	if i != -1 {
		t.Errorf("CrystalBall = %d; expected -1", i)
	}

	b = []bool{false, false, false, false, false, false, false, false, false, false, false, true}
	i = CrystalBallTwo(b)
	if i != 11 {
		t.Errorf("CrystalBall = %d, expected 11", i)
	}

	b = []bool{}
	i = CrystalBallTwo(b)
	if i != -1 {
		t.Errorf("CrystalBall = %d, expected -1", i)
	}
}

func TestBlockSearch(t *testing.T) {
	a := []int{0, 1, 5, 6, 9, 12, 22, 36, 45, 57, 78, 101, 120, 125, 144, 150, 161, 181, 186, 190, 200}
	v := 144
	i := BlockSearch(a, v)
	if i != 14 {
		t.Errorf("BlockSearch = %d, expected 14", i)
	}

	v = 22
	i = BlockSearch(a, v)
	if i != 6 {
		t.Errorf("BlockSearch = %d, expected 6", i)
	}

	a = []int{0, 2, 5, 9, 10, 13, 16, 17, 18, 25, 26, 30}
	v = 2
	i = BlockSearch(a, v)
	if i != 1 {
		t.Errorf("BlockSearch = %d, expected 1", i)
	}

	v = 18
	i = BlockSearch(a, v)
	if i != 8 {
		t.Errorf("BlockSearch = %d, expected 8", i)
	}
}

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
