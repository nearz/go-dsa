package graph

import (
	"fmt"
	"testing"
)

func TestAddEdge(t *testing.T) {
	g := New[int]()
	v0 := NewVertex(0)
	v1 := NewVertex(1)
	v2 := NewVertex(2)
	v3 := NewVertex(3)
	v4 := NewVertex(4)
	g.AddEdge(v0, v1, 1)
	g.AddEdge(v0, v2, 5)
	g.AddEdge(v1, v2, 7)
	g.AddEdge(v1, v3, 3)
	g.AddEdge(v3, v2, 2)
	g.AddEdge(v2, v4, 1)
	g.PrintVerticies()
	fmt.Printf("Size: %d\n", g.Size())
}

func TestHasEdge(t *testing.T) {
	g := New[int]()
	v0 := NewVertex(0)
	v1 := NewVertex(1)
	v2 := NewVertex(2)
	v3 := NewVertex(3)
	v4 := NewVertex(4)
	g.AddEdge(v0, v1, 1)
	g.AddEdge(v0, v2, 5)
	g.AddEdge(v1, v2, 7)
	g.AddEdge(v1, v3, 3)
	g.AddEdge(v3, v2, 2)
	g.AddEdge(v2, v4, 1)
	g.PrintVerticies()
	fmt.Printf("Has Edge: %t", g.HasEdge(v3, v2))
}

func TestRemoveEdge(t *testing.T) {
	g := New[int]()
	v0 := NewVertex(0)
	v1 := NewVertex(1)
	v2 := NewVertex(2)
	v3 := NewVertex(3)
	v4 := NewVertex(4)
	g.AddEdge(v0, v1, 1)
	g.AddEdge(v0, v2, 5)
	g.AddEdge(v1, v2, 7)
	g.AddEdge(v1, v3, 3)
	g.AddEdge(v3, v2, 2)
	g.AddEdge(v2, v4, 1)
	g.PrintVerticies()
	fmt.Printf("Removed Edge: %t", g.RemoveEdge(v4, v3))
	fmt.Println("")
	g.PrintVerticies()
}

func TestNeighbors(t *testing.T) {
	g := New[int]()
	v0 := NewVertex(0)
	v1 := NewVertex(1)
	v2 := NewVertex(2)
	v3 := NewVertex(3)
	v4 := NewVertex(4)
	g.AddEdge(v0, v1, 1)
	g.AddEdge(v0, v2, 5)
	g.AddEdge(v1, v2, 7)
	g.AddEdge(v1, v3, 3)
	g.AddEdge(v3, v2, 2)
	g.AddEdge(v2, v4, 1)
	g.PrintVerticies()
	n := g.Neighbors(v0)
	PrintEdges(n)
}

func TestBFS(t *testing.T) {
	g := New[int]()
	v0 := NewVertex(0)
	v1 := NewVertex(1)
	v2 := NewVertex(2)
	v3 := NewVertex(3)
	v4 := NewVertex(4)
	g.AddEdge(v0, v1, 1)
	g.AddEdge(v0, v2, 5)
	g.AddEdge(v1, v2, 7)
	g.AddEdge(v1, v3, 3)
	g.AddEdge(v3, v2, 2)
	g.AddEdge(v2, v4, 1)
	g.PrintVerticies()
	ok, vrts, err := g.BFS(v0, v4)
	if err != nil {
		fmt.Println(err)
		return
	}
	if ok {
		PrintVertexSlice(vrts)
	} else {
		fmt.Println("Path not found")
	}
}

func TestDSP(t *testing.T) {
	g := New[int]()
	v0 := NewVertex(0)
	v1 := NewVertex(1)
	v2 := NewVertex(2)
	v3 := NewVertex(3)
	v4 := NewVertex(4)
	g.AddEdge(v0, v1, 1)
	g.AddEdge(v0, v2, 5)
	g.AddEdge(v1, v2, 1)
	g.AddEdge(v1, v3, 3)
	g.AddEdge(v3, v2, 2)
	g.AddEdge(v2, v4, 1)
	g.PrintVerticies()
	ok, vrts := g.DSP(v0, v4)
	if ok {
		PrintVertexSlice(vrts)
	} else {
		fmt.Println("Path not found")
	}
}
