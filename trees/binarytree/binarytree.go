package binarytree

import (
	"github.com/nearz/go-dsa/queues/queue"
)

// TODO: Add Comments
// TODO: Do a balanced insert?

type BinaryNode[T comparable] struct {
	val   T
	left  *BinaryNode[T]
	right *BinaryNode[T]
}

func New[T comparable](val T) *BinaryNode[T] {
	return &BinaryNode[T]{
		val:   val,
		left:  nil,
		right: nil,
	}
}

func (b *BinaryNode[T]) GetValue() T {
	return b.val
}

func (b *BinaryNode[T]) PreOrder() []T {
	path := []T{}
	path = preOrder(b, path)
	return path
}

func preOrder[T comparable](b *BinaryNode[T], path []T) []T {
	if b == nil {
		return path
	}
	path = append(path, b.val)
	path = preOrder(b.left, path)
	path = preOrder(b.right, path)

	return path
}

func (b *BinaryNode[T]) InOrder() []T {
	path := []T{}
	path = inOrder(b, path)
	return path
}

func inOrder[T comparable](b *BinaryNode[T], path []T) []T {
	if b == nil {
		return path
	}
	path = inOrder(b.left, path)
	path = append(path, b.val)
	path = inOrder(b.right, path)

	return path
}

func (b *BinaryNode[T]) PostOrder() []T {
	path := []T{}
	path = postOrder(b, path)
	return path
}

func postOrder[T comparable](b *BinaryNode[T], path []T) []T {
	if b == nil {
		return path
	}
	path = postOrder(b.left, path)
	path = postOrder(b.right, path)
	path = append(path, b.val)

	return path
}

func (b *BinaryNode[T]) BFSSearch(target T) (bool, *BinaryNode[T], error) {
	q := queue.New[*BinaryNode[T]]()
	q.Enqueue(b)
	for !q.Empty() {
		curr, err := q.Deque()
		if err != nil {
			return false, nil, err
		}
		if curr.val == target {
			return true, curr, nil
		}
		q.Enqueue(curr.left)
		q.Enqueue(curr.right)
	}
	return false, nil, nil
}
