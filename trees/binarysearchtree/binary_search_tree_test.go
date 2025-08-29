package binarysearchtree

import (
	"slices"
	"testing"
)

func TestSearch(t *testing.T) {
	bst := New[int]()
	bst.Insert(12)
	bst.Insert(16)
	bst.Insert(14)
	bst.Insert(17)
	bst.Insert(7)
	bst.Insert(5)
	bst.Insert(8)
	target := 14
	ok, _ := bst.Search(target)
	if !ok {
		t.Errorf("BST Search: value not found, %v\n", target)
		return
	}
}

func TestSize(t *testing.T) {
	bst := New[int]()
	bst.Insert(12)
	bst.Insert(16)
	bst.Insert(5)
	v := 3
	s := bst.Size()
	if s != v {
		t.Errorf("Size error, res: %d, expected: %d", s, v)
	}

	bst.Insert(12)
	bst.Insert(18)
	bst.Insert(1)
	v = 6
	s = bst.Size()
	if s != v {
		t.Errorf("Size error, res: %d, expected: %d", s, v)
	}
}

func TestSizeWithDuplicate(t *testing.T) {
	bst := New[int]()
	bst.Insert(12)
	bst.Insert(16)
	bst.Insert(5)
	v := 3
	s := bst.Size()
	if s != v {
		t.Errorf("Size error, res: %d, expected: %d", s, v)
	}

	bst.Insert(12)
	v = 4
	s = bst.Size()
	if s != v {
		t.Errorf("Size error, res: %d, expected: %d", s, v)
	}
}

func TestSizeAfterDelete(t *testing.T) {
	bst := New[int]()
	bst.Insert(12)
	bst.Insert(16)
	bst.Insert(5)
	v := 3
	s := bst.Size()
	if s != v {
		t.Errorf("Size error, res: %d, expected: %d", s, v)
	}

	bst.Delete(16)
	v = 2
	s = bst.Size()
	if s != v {
		t.Errorf("Size error, res: %d, expected: %d", s, v)
	}

	bst.Delete(5)
	v = 1
	s = bst.Size()
	if s != v {
		t.Errorf("Size error, res: %d, expected: %d", s, v)
	}

	bst.Delete(12)
	v = 0
	s = bst.Size()
	if s != v {
		t.Errorf("Size error, res: %d, expected: %d", s, v)
	}
}

func TestSearchAboveMax(t *testing.T) {
	bst := New[int]()
	bst.Insert(12)
	bst.Insert(16)
	bst.Insert(14)
	bst.Insert(17)
	bst.Insert(7)
	bst.Insert(5)
	bst.Insert(8)
	target := 18
	ok, _ := bst.Search(target)
	if ok {
		t.Errorf("BST Search: value not found, %v\n", target)
		return
	}
}

func TestSearchBelowMin(t *testing.T) {
	bst := New[int]()
	bst.Insert(12)
	bst.Insert(16)
	bst.Insert(14)
	bst.Insert(17)
	bst.Insert(7)
	bst.Insert(5)
	bst.Insert(8)
	target := 3
	ok, _ := bst.Search(target)
	if ok {
		t.Errorf("BST Search: value not found, %v\n", target)
		return
	}
}

func TestInsert(t *testing.T) {
	bst := New[int]()
	bst.Insert(12)
	bst.Insert(16)
	bst.Insert(14)
	bst.Insert(17)
	bst.Insert(7)
	bst.Insert(5)
	bst.Insert(8)
	val := 9
	bst.Insert(val)
	s := []int{12, 7, 5, 8, 9, 16, 14, 17}
	path := bst.PreOrderPath()
	if !slices.Equal(s, path) {
		t.Errorf("Insert error: res: %v, expected: %v\n", path, s)
	}
}

func TestInsertAndMin(t *testing.T) {
	bst := New[int]()
	bst.Insert(12)
	bst.Insert(16)
	bst.Insert(14)
	bst.Insert(17)
	bst.Insert(7)
	bst.Insert(5)
	bst.Insert(8)
	m := bst.Min()
	v := 5
	if m != v {
		t.Errorf("Min error: res: %d, expected: %d", m, v)
	}
}

func TestInsertAndMax(t *testing.T) {
	bst := New[int]()
	bst.Insert(12)
	bst.Insert(16)
	bst.Insert(14)
	bst.Insert(17)
	bst.Insert(7)
	bst.Insert(5)
	bst.Insert(8)
	m := bst.Max()
	v := 17
	if m != v {
		t.Errorf("Max error: res: %d, expected: %d", m, v)
	}
}

func TestDeleteMax(t *testing.T) {
	bst := New[int]()
	bst.Insert(15)
	bst.Insert(7)
	bst.Insert(4)
	bst.Insert(51)
	bst.Insert(25)
	bst.Insert(37)
	bst.Insert(32)
	bst.Insert(99)
	bst.Insert(98)
	bst.Insert(101)
	s := []int{15, 7, 4, 51, 25, 37, 32, 99, 98}
	val := 101
	bst.Delete(val)
	path := bst.PreOrderPath()
	if !slices.Equal(path, s) {
		t.Errorf("Delete Error: res: %v, expected: %v\n", path, s)
	}
	if bst.Max() != 99 {
		t.Errorf("Delete and new min error: res: %d, expected: %d", bst.Min(), 99)
	}
}

func TestDeleteMin(t *testing.T) {
	bst := New[int]()
	bst.Insert(15)
	bst.Insert(7)
	bst.Insert(4)
	bst.Insert(51)
	bst.Insert(25)
	bst.Insert(37)
	bst.Insert(32)
	bst.Insert(99)
	bst.Insert(98)
	bst.Insert(101)
	s := []int{15, 7, 51, 25, 37, 32, 99, 98, 101}
	val := 4
	bst.Delete(val)
	path := bst.PreOrderPath()
	if !slices.Equal(path, s) {
		t.Errorf("Delete Error: res: %v, expected: %v\n", path, s)
	}
	if bst.Min() != 7 {
		t.Errorf("Delete and new min error: res: %d, expected: %d", bst.Min(), 7)
	}
}

func TestDeleteWithSubtree(t *testing.T) {
	bst := New[int]()
	bst.Insert(15)
	bst.Insert(7)
	bst.Insert(4)
	bst.Insert(51)
	bst.Insert(25)
	bst.Insert(37)
	bst.Insert(32)
	bst.Insert(99)
	bst.Insert(98)
	bst.Insert(101)
	s := []int{15, 7, 4, 98, 25, 37, 32, 99, 101}
	val := 51
	bst.Delete(val)
	path := bst.PreOrderPath()
	if !slices.Equal(path, s) {
		t.Errorf("Delete Error: res: %v, expected: %v\n", path, s)
	}
}

func TestDeleteShiftValCount(t *testing.T) {
	bst := New[int]()
	bst.Insert(15)
	bst.Insert(7)
	bst.Insert(4)
	bst.Insert(51)
	bst.Insert(25)
	bst.Insert(37)
	bst.Insert(32)
	bst.Insert(99)
	bst.Insert(98)
	bst.Insert(98)
	bst.Insert(101)
	s := []int{15, 7, 4, 98, 98, 25, 37, 32, 99, 101}
	val := 51
	bst.Delete(val)
	path := bst.PreOrderPath()
	if !slices.Equal(path, s) {
		t.Errorf("Delete Error: res: %v, expected: %v\n", path, s)
	}
}

func TestDeleteNonexistentValue(t *testing.T) {
	bst := New[int]()
	bst.Insert(15)
	bst.Insert(7)
	bst.Insert(4)
	bst.Insert(51)
	bst.Insert(25)
	bst.Insert(37)
	bst.Insert(32)
	bst.Insert(99)
	bst.Insert(98)
	bst.Insert(101)
	s := []int{15, 7, 4, 51, 25, 37, 32, 99, 98, 101}
	val := 31
	bst.Delete(val)
	path := bst.PreOrderPath()
	if !slices.Equal(path, s) {
		t.Errorf("Delete Error: res: %v, expected: %v\n", path, s)
	}
}

func TestDeleteOneNode(t *testing.T) {
	bst := New[int]()
	bst.Insert(15)
	s := []int{}
	val := 15
	bst.Delete(val)
	path := bst.PreOrderPath()
	if !slices.Equal(path, s) {
		t.Errorf("Delete Error: res: %v, expected: %v\n", path, s)
	}
	if bst.Size() != 0 {
		t.Errorf("Delete one node error: size: %d, expected: %d", bst.Size(), 0)
	}
	if bst.Max() != 0 {
		t.Errorf("Delete one node error: max %d, expected: %d", bst.Max(), 0)
	}
	if bst.Min() != 0 {
		t.Errorf("Delete one node error: min %d, expected: %d", bst.Min(), 0)
	}
	if bst.GetHeight() != -1 {
		t.Errorf("Delete one node error: height %d, expected: %d", bst.Min(), -1)
	}
}

func TestDeleteOneNodeInsertNewRoot(t *testing.T) {
	bst := New[int]()
	bst.Insert(15)
	s := []int{}
	val := 15
	bst.Delete(val)
	path := bst.PreOrderPath()
	if !slices.Equal(path, s) {
		t.Errorf("Delete Error: res: %v, expected: %v\n", path, s)
	}
	if bst.Size() != 0 {
		t.Errorf("Delete one node error: size: %d, expected: %d", bst.Size(), 0)
	}
	if bst.Max() != 0 {
		t.Errorf("Delete one node error: max %d, expected: %d", bst.Max(), 0)
	}
	if bst.Min() != 0 {
		t.Errorf("Delete one node error: min %d, expected: %d", bst.Min(), 0)
	}
	if bst.GetHeight() != -1 {
		t.Errorf("Delete one node error: height %d, expected: %d", bst.Min(), -1)
	}

	val = 4
	bst.Insert(val)
	s = []int{4}
	path = bst.PreOrderPath()
	if !slices.Equal(path, s) {
		t.Errorf("Delete Error: res: %v, expected: %v\n", path, s)
	}
	if bst.Size() != 1 {
		t.Errorf("Delete root and insert: size: %d, expected: %d", bst.Size(), 1)
	}
	if bst.Max() != val {
		t.Errorf("Delete root and insert: max %d, expected: %d", bst.Max(), val)
	}
	if bst.Min() != val {
		t.Errorf("Delete root and insert: min %d, expected: %d", bst.Min(), val)
	}
	if bst.GetHeight() != 0 {
		t.Errorf("Delete root and insert: height %d, expected: %d", bst.Min(), 0)
	}
	if bst.root.height != 0 {
		t.Errorf("Delete root and insert error: root height: %d, expected: %d", bst.root.height, 0)
	}
}

func TestDeleteLeaf(t *testing.T) {
	bst := New[int]()
	bst.Insert(15)
	bst.Insert(7)
	bst.Insert(4)
	bst.Insert(51)
	bst.Insert(25)
	bst.Insert(37)
	bst.Insert(32)
	bst.Insert(99)
	bst.Insert(98)
	bst.Insert(101)
	s := []int{15, 7, 51, 25, 37, 32, 99, 98, 101}
	val := 4
	bst.Delete(val)
	path := bst.PreOrderPath()
	if !slices.Equal(path, s) {
		t.Errorf("Delete Error: res: %v, expected: %v\n", path, s)
	}
}

func TestDeleteOneChild(t *testing.T) {
	bst := New[int]()
	bst.Insert(15)
	bst.Insert(7)
	bst.Insert(4)
	bst.Insert(51)
	bst.Insert(25)
	bst.Insert(37)
	bst.Insert(32)
	bst.Insert(99)
	bst.Insert(98)
	bst.Insert(101)
	s := []int{15, 4, 51, 25, 37, 32, 99, 98, 101}
	val := 7
	bst.Delete(val)
	path := bst.PreOrderPath()
	if !slices.Equal(path, s) {
		t.Errorf("Delete Error: res: %v, expected: %v\n", path, s)
	}
}

func TestDeleteOneChildTwo(t *testing.T) {
	bst := New[int]()
	bst.Insert(15)
	bst.Insert(7)
	bst.Insert(4)
	bst.Insert(51)
	bst.Insert(25)
	bst.Insert(37)
	bst.Insert(32)
	bst.Insert(99)
	bst.Insert(98)
	bst.Insert(101)
	s := []int{15, 7, 4, 51, 37, 32, 99, 98, 101}
	val := 25
	bst.Delete(val)
	path := bst.PreOrderPath()
	if !slices.Equal(path, s) {
		t.Errorf("Delete Error: res: %v, expected: %v\n", path, s)
	}
}

func TestHeightAfterDelete(t *testing.T) {
	bst := New[int]()
	bst.Insert(15)
	bst.Insert(7)
	bst.Insert(4)
	bst.Insert(51)
	bst.Insert(25)
	bst.Insert(37)
	bst.Insert(32)
	bst.Insert(99)
	bst.Insert(98)
	bst.Insert(101)
	val := 32
	bst.Delete(val)
	if ok, n := bst.Search(37); ok {
		h := n.GetNodeHeight()
		newH := 0
		if h != newH {
			t.Errorf("Node height not updated on delete, res: %d, expected: %d\n", h, newH)
		}
	}
}

func TestHeightAfterDeleteTwo(t *testing.T) {
	bst := New[int]()
	bst.Insert(15)
	bst.Insert(7)
	bst.Insert(4)
	bst.Insert(51)
	bst.Insert(25)
	bst.Insert(99)
	bst.Insert(98)
	val := 51
	bst.Delete(val)
	if ok, n := bst.Search(98); ok {
		h := n.GetNodeHeight()
		newH := 1
		if h != newH {
			t.Errorf("Node height not updated on delete, res: %d, expected: %d\n", h, newH)
		}
	}
}

func TestHeightAfterDeleteThree(t *testing.T) {
	bst := New[int]()
	bst.Insert(15)
	bst.Insert(7)
	bst.Insert(4)
	bst.Insert(51)
	bst.Insert(25)
	bst.Insert(99)
	bst.Insert(98)
	val := 15
	bst.Delete(val)
	if ok, n := bst.Search(25); ok {
		h := n.GetNodeHeight()
		newH := 3
		if h != newH {
			t.Errorf("Node height not updated on delete, res: %d, expected: %d\n", h, newH)
		}
	}
}

func TestNodeHeight(t *testing.T) {
	bst := New[int]()
	bst.Insert(12)
	bst.Insert(16)
	bst.Insert(14)
	bst.Insert(17)
	bst.Insert(7)
	bst.Insert(5)
	bst.Insert(8)
	bst.Insert(13)
	if ok, node := bst.Search(16); ok {
		h := node.GetNodeHeight()
		if h != 2 {
			t.Errorf("Node height error: res: %d, expected: %d\n", h, 2)
		}
	}
}

func TestNodeHeight2(t *testing.T) {
	bst := New[int]()
	bst.Insert(12)
	bst.Insert(16)
	bst.Insert(14)
	bst.Insert(17)
	bst.Insert(30)
	bst.Insert(35)
	bst.Insert(25)
	bst.Insert(7)
	bst.Insert(5)
	bst.Insert(8)
	bst.Insert(13)
	if ok, node := bst.Search(16); ok {
		h := node.GetNodeHeight()
		if h != 3 {
			t.Errorf("Node height error: res: %d, expected: %d\n", h, 3)
		}
	}

	if ok, node := bst.Search(14); ok {
		h := node.GetNodeHeight()
		if h != 1 {
			t.Errorf("Node height error: res: %d, expected: %d\n", h, 1)
		}
	}
}
