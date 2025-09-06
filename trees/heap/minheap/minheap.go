package minheap

import (
	"cmp"
	"errors"
)

type MinHeap[T cmp.Ordered] struct {
	heap []T
}

func New[T cmp.Ordered]() *MinHeap[T] {
	return &MinHeap[T]{heap: []T{}}
}

func (m *MinHeap[T]) Insert(val T) {
	m.heap = append(m.heap, val)
	m.heapifyUp(len(m.heap) - 1)
}

func (m *MinHeap[T]) heapifyUp(idx int) {
	if idx == 0 {
		return
	}
	pi := parent(idx)
	parentValue := m.heap[pi]
	curr := m.heap[idx]
	if curr < parentValue {
		m.heap[pi] = curr
		m.heap[idx] = parentValue
		m.heapifyUp(pi)
	}
}

func (m *MinHeap[T]) Pop() (T, error) {
	if m.Empty() {
		var z T
		return z, errors.New("Empty MinHeap")
	}
	value := m.heap[0]
	if len(m.heap) == 1 {
		m.heap = []T{}
		return value, nil
	}
	m.heap[0] = m.heap[len(m.heap)-1]
	m.heap = m.heap[:len(m.heap)-1]
	m.heapifyDown(0)
	return value, nil
}

func (m *MinHeap[T]) heapifyDown(idx int) {
	n := len(m.heap)
	li := leftChild(idx)
	if li >= n {
		return
	}
	ri := rightChild(idx)
	small := li
	if ri < n && m.heap[ri] < m.heap[li] {
		small = ri
	}
	if m.heap[idx] <= m.heap[small] {
		return
	}
	m.heap[small], m.heap[idx] = m.heap[idx], m.heap[small]
	m.heapifyDown(small)
}

func (m *MinHeap[T]) Peek() (T, error) {
	if m.Empty() {
		var z T
		return z, errors.New("Empty MinHeap")
	}
	return m.heap[0], nil
}

func (m *MinHeap[T]) Empty() bool {
	if len(m.heap) == 0 {
		return true
	}
	return false
}

func parent(i int) int {
	return (i - 1) / 2
}

func leftChild(i int) int {
	return 2*i + 1
}

func rightChild(i int) int {
	return 2*i + 2
}
