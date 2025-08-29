package binarysearchtree

import (
	"cmp"
)

// TODO: Add Comments
// TODO: Iterator support?

// Binary Node
type BinaryNode[T cmp.Ordered] struct {
	val      T
	valCount int
	height   int
	left     *BinaryNode[T]
	right    *BinaryNode[T]
}

func (b *BinaryNode[T]) GetValue() T {
	return b.val
}

func (b *BinaryNode[T]) GetNodeHeight() int {
	if b == nil {
		return -1
	}
	return b.height
}

// Binary Search Tree
type BinarySearchTree[T cmp.Ordered] struct {
	root   *BinaryNode[T]
	height int
	size   int
	min    T
	max    T
}

func New[T cmp.Ordered]() *BinarySearchTree[T] {
	b := &BinarySearchTree[T]{root: nil, height: 0}
	return b
}

func (b *BinarySearchTree[T]) GetHeight() int {
	if b == nil || b.root == nil {
		return -1
	}
	return b.height
}

func (b *BinarySearchTree[T]) Max() T {
	if b == nil || b.root == nil {
		var z T
		return z
	}
	return b.max
}

func (b *BinarySearchTree[T]) Min() T {
	if b == nil || b.root == nil {
		var z T
		return z
	}
	return b.min
}

func (b *BinarySearchTree[T]) Size() int {
	if b == nil || b.root == nil {
		return 0
	}
	return b.size
}

// Search
func (b *BinarySearchTree[T]) Search(target T) (bool, *BinaryNode[T]) {
	if b == nil || b.root == nil {
		return false, nil
	}
	if b.root != nil && target > b.max {
		return false, nil
	}
	if b.root != nil && target < b.min {
		return false, nil
	}
	if b.root.val == target {
		return true, b.root
	}
	return search(b.root, target)
}

func search[T cmp.Ordered](b *BinaryNode[T], target T) (bool, *BinaryNode[T]) {
	if b == nil {
		return false, nil
	}
	if b.val == target {
		return true, b
	}

	if b.val < target {
		return search(b.right, target)
	}
	return search(b.left, target)
}

// Insert
func (b *BinarySearchTree[T]) Insert(val T) {
	if b.root == nil {
		b.root = &BinaryNode[T]{val: val, valCount: 1, height: 0}
		b.height = b.root.height
		b.min = val
		b.max = val
		b.size++
		return
	}
	b.min = min(b.min, val)
	b.max = max(b.max, val)
	insert(b.root, val)
	b.height = b.root.height
	b.size++
}

func insert[T cmp.Ordered](b *BinaryNode[T], val T) {
	if b.val == val {
		b.valCount++
		return
	}
	if val < b.val {
		if b.left == nil {
			b.left = &BinaryNode[T]{val: val, valCount: 1, height: 0}
		} else {
			insert(b.left, val)
		}
	} else {
		if b.right == nil {
			b.right = &BinaryNode[T]{val: val, valCount: 1, height: 0}
		} else {
			insert(b.right, val)
		}
	}

	maxChildHeight := max(b.left.GetNodeHeight(), b.right.GetNodeHeight())
	b.height = maxChildHeight + 1
}

// Delete
func (b *BinarySearchTree[T]) Delete(target T) *BinaryNode[T] {
	if b == nil || b.root == nil {
		return nil
	}
	if b.root.left == nil && b.root.right == nil && b.root.val == target {
		b.root = nil
		b.size = 0
		b.height = -1
		var z T
		b.min = z
		b.max = z
		return nil
	}
	var ok bool
	b.root, ok = delete(b.root, target)
	if ok {
		b.size--
	}
	b.height = b.root.height
	if target == b.min {
		b.min = findMin(b.root)
	} else if target == b.max {
		b.max = findMax(b.root)
	}
	return b.root
}

func findMin[T cmp.Ordered](b *BinaryNode[T]) T {
	if b == nil {
		var z T
		return z
	}
	curr := b
	for curr.left != nil {
		curr = curr.left
	}
	return curr.val
}

func findMax[T cmp.Ordered](b *BinaryNode[T]) T {
	if b == nil {
		var z T
		return z
	}
	curr := b
	for curr.right != nil {
		curr = curr.right
	}
	return curr.val
}

func delete[T cmp.Ordered](b *BinaryNode[T], target T) (*BinaryNode[T], bool) {
	var ok bool
	if b == nil {
		return nil, false
	}
	if target < b.val {
		b.left, ok = delete(b.left, target)
	} else if target > b.val {
		b.right, ok = delete(b.right, target)
	} else {
		if b.valCount > 1 {
			b.valCount--
			return b, true
		}
		if b.left == nil {
			t := b.right
			b = nil
			return t, true
		}
		if b.right == nil {
			t := b.left
			b = nil
			return t, true
		}
		curr := b.right
		for curr.left != nil {
			curr = curr.left
		}
		b.val = curr.val
		b.valCount = curr.valCount
		curr.valCount = 1
		b.right, ok = delete(b.right, curr.val)
	}
	maxChildHeight := max(b.left.GetNodeHeight(), b.right.GetNodeHeight())
	b.height = maxChildHeight + 1
	return b, ok
}

// Utilities
func (b *BinarySearchTree[T]) PreOrderPath() []T {
	path := &[]T{}
	if b.root == nil {
		return *path
	}
	preOrderPath(b.root, path)
	return *path
}

func preOrderPath[T cmp.Ordered](b *BinaryNode[T], path *[]T) *[]T {
	if b == nil {
		return path
	}

	for range b.valCount {
		*path = append(*path, b.val)
	}
	preOrderPath(b.left, path)
	preOrderPath(b.right, path)

	return path
}
