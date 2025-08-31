package minheap

import (
	"slices"
	"testing"
)

func TestHeap(t *testing.T) {
	h := New[int]()
	h.Insert(50)
	h.Insert(71)
	h.Insert(100)
	h.Insert(101)
	h.Insert(80)
	h.Insert(200)
	h.Insert(101)
	h.Insert(3)
	s := []int{3, 50, 100, 71, 80, 200, 101, 101}
	if !slices.Equal(s, h.heap) {
		t.Errorf("Insert error: res %v, expected: %v", h.heap, s)
	}
}

func TestHeapOneInsert(t *testing.T) {
	h := New[int]()
	h.Insert(1)
	s := []int{1}
	if !slices.Equal(s, h.heap) {
		t.Errorf("Insert error: res %v, expected: %v", h.heap, s)
	}
}

func TestHeapTwo(t *testing.T) {
	h := New[int]()
	h.Insert(50)
	h.Insert(55)
	h.Insert(22)
	h.Insert(60)
	h.Insert(45)
	h.Insert(33)
	h.Insert(61)
	h.Insert(59)
	h.Insert(28)
	h.Insert(39)
	h.Insert(51)
	s := []int{22, 28, 33, 45, 39, 50, 61, 60, 59, 55, 51}
	if !slices.Equal(s, h.heap) {
		t.Errorf("Insert error: res %v, expected: %v", h.heap, s)
	}
}

func TestHeapThree(t *testing.T) {
	h := New[int]()
	h.Insert(7)
	h.Insert(5)
	h.Insert(6)
	h.Insert(4)
	s := []int{4, 5, 6, 7}
	if !slices.Equal(s, h.heap) {
		t.Errorf("Insert error: res %v, expected: %v", h.heap, s)
	}
}

func TestPop(t *testing.T) {
	h := New[int]()
	h.Insert(50)
	h.Insert(71)
	h.Insert(100)
	h.Insert(101)
	h.Insert(80)
	h.Insert(200)
	h.Insert(3)
	p, err := h.Pop()
	s := []int{50, 71, 100, 101, 80, 200}
	if err != nil {
		t.Error("Pop error: expected a nil error")
	}
	if p != 3 {
		t.Errorf("Pop error: res: %d, expected: %d", p, 3)
	}
	if !slices.Equal(s, h.heap) {
		t.Errorf("Pop error not heapified, res: %v, expected: %v", h.heap, s)
	}
}

func TestPopThree(t *testing.T) {
	h := New[int]()
	h.Insert(4)
	h.Insert(5)
	h.Insert(6)
	h.Insert(7)
	p, err := h.Pop()
	s := []int{5, 7, 6}
	if err != nil {
		t.Error("Pop error: expected a nil error")
	}
	if p != 4 {
		t.Errorf("Pop error: res: %d, expected: %d", p, 3)
	}
	if !slices.Equal(s, h.heap) {
		t.Errorf("Pop error not heapified, res: %v, expected: %v", h.heap, s)
	}
}

func TestPopTwo(t *testing.T) {
	h := New[int]()
	h.Insert(50)
	h.Insert(55)
	h.Insert(22)
	h.Insert(60)
	h.Insert(45)
	h.Insert(33)
	h.Insert(61)
	h.Insert(59)
	h.Insert(28)
	h.Insert(39)
	h.Insert(51)
	p, err := h.Pop()
	s := []int{28, 39, 33, 45, 51, 50, 61, 60, 59, 55}
	if err != nil {
		t.Error("Pop error: expected a nil error")
	}
	if p != 22 {
		t.Errorf("Pop error: res: %d, expected: %d", p, 3)
	}
	if !slices.Equal(s, h.heap) {
		t.Errorf("Pop error not heapified, res: %v, expected: %v", h.heap, s)
	}
}

func TestPeek(t *testing.T) {
	h := New[int]()
	h.Insert(50)
	h.Insert(71)
	h.Insert(100)
	h.Insert(101)
	h.Insert(80)
	h.Insert(200)
	h.Insert(101)
	h.Insert(3)
	v := 3
	p, err := h.Peek()
	if err != nil {
		t.Error("Peek error: expected no error")
	}
	if v != p {
		t.Errorf("Peek error: res: %d, expected: %d", p, v)
	}
}

func TestPeekEmpty(t *testing.T) {
	h := New[int]()
	_, err := h.Peek()
	if err == nil {
		t.Error("Expected peek error on empty minheap")
	}
}

func TestEmpty(t *testing.T) {
	h := New[int]()
	b := h.Empty()
	if !b {
		t.Error("Empty error: expected true")
	}

	h.Insert(1)
	b = h.Empty()
	if b {
		t.Error("Emptry error: expected false")
	}
}
