package binarytree

import (
	"fmt"
	"testing"
)

// func TestPreOrder(t *testing.T) {
// 	root := New(8)
// 	root.left = New(7)
// 	root.right = New(6)
// 	root.left.left = New(5)
// 	root.left.right = New(4)
// 	path := root.PreOrder()
// 	fmt.Println(path)
// }
//
// func TestInOrder(t *testing.T) {
// 	root := New(8)
// 	root.left = New(7)
// 	root.right = New(6)
// 	root.left.left = New(5)
// 	root.left.right = New(4)
// 	path := root.InOrder()
// 	fmt.Println(path)
// }
//
// func TestPostOrder(t *testing.T) {
// 	root := New(8)
// 	root.left = New(7)
// 	root.right = New(6)
// 	root.left.left = New(5)
// 	root.left.right = New(4)
// 	path := root.PostOrder()
// 	fmt.Println(path)
// }

func TestBFSSearch(t *testing.T) {
	root := New(8)
	root.left = New(7)
	root.right = New(6)
	root.left.left = New(5)
	root.left.right = New(4)
	ok, node, err := root.BFSSearch(4)
	if err != nil {
		fmt.Printf("BFS Search Err: %v", err)
		return
	}
	if !ok {
		fmt.Printf("BSF Search: node note found")
		return
	}
	fmt.Printf("Node Found: value = %v\n", node.GetValue())
}
