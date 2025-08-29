package binarytree

func EqualBinaryTree[T comparable](r1 *BinaryNode[T], r2 *BinaryNode[T]) bool {
	if r1 == nil && r2 == nil {
		return true
	}
	if r1 == nil || r2 == nil {
		return false
	}
	if r1.val != r2.val {
		return false
	}

	return EqualBinaryTree(r1.left, r2.left) && EqualBinaryTree(r1.right, r2.right)
}
