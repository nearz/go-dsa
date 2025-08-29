package binarytree

import (
	"fmt"
	"testing"
)

func TestPreOrder(t *testing.T) {
	root := Create(8, 7, 6, 5, 4)
	path := *root.PreOrder()
	fmt.Println(path)
}

func TestInOrder(t *testing.T) {
	root := Create(8, 7, 6, 5, 4)
	path := *root.InOrder()
	fmt.Println(path)
}

func TestPostOrder(t *testing.T) {
	root := Create(8, 7, 6, 5, 4)
	path := *root.PostOrder()
	fmt.Println(path)
}

func TestEqualBinaryTree(t *testing.T) {
	r1 := Create(8, 7, 6, 5, 4)
	r2 := Create(8, 7, 5, 4)
	res := EqualBinaryTree(r1, r2)
	fmt.Println(res)
}

func TestBFSSearch(t *testing.T) {
	root := Create(8, 7, 6, 5, 4)
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

func TestDFSSearch(t *testing.T) {
	root := Create(8, 7, 6, 5, 4)
	ok, node := root.DFSSearch(4)
	if !ok {
		fmt.Printf("BSF Search: node note found")
		return
	}
	fmt.Printf("Node Found: value = %v\n", node.GetValue())
}
