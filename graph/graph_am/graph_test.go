package graph_am

import (
	"fmt"
	"testing"
)

// func TestNewGraph(t *testing.T) {
// 	g, err := NewGraph(20)
// 	if err != nil {
// 		t.Errorf("Error: %v", err)
// 	}
// 	g.Print()
// }

// func TestGraphAddEdge(t *testing.T) {
// 	g, err := NewGraph(5)
// 	if err != nil {
// 		t.Errorf("Error: %v", err)
// 	}
// 	g.AddEdge(0, 1)
// 	g.AddEdge(0, 2)
// 	g.AddEdge(0, 3)
// 	g.AddEdge(1, 3)
// 	g.AddEdge(3, 4)
// 	g.Print()
// }

// func TestGraphRemoveEdge(t *testing.T) {
// 	g, err := NewGraph(5)
// 	if err != nil {
// 		t.Errorf("Error: %v", err)
// 	}
// 	g.AddEdge(0, 1)
// 	g.AddEdge(0, 2)
// 	g.AddEdge(0, 3)
// 	g.AddEdge(1, 3)
// 	g.AddEdge(3, 4)
// 	g.Print()
// 	res := g.RemoveEdge(0, 3)
// 	fmt.Println(res)
// 	fmt.Println("")
// 	g.Print()
// }

// func TestGraphHasEdge(t *testing.T) {
// 	g, err := NewGraph(5)
// 	if err != nil {
// 		t.Errorf("Error: %v", err)
// 	}
// 	g.AddEdge(0, 1)
// 	g.AddEdge(0, 2)
// 	g.AddEdge(0, 3)
// 	g.AddEdge(1, 3)
// 	g.AddEdge(3, 4)
// 	g.Print()
// 	res := g.HasEdge(0, 3)
// 	if res {
// 		fmt.Println("Edge exists")
// 	} else {
// 		fmt.Println("Edge not found")
// 	}
// }

// func TestGraphNeighbors(t *testing.T) {
// 	g, err := NewGraph(5)
// 	if err != nil {
// 		t.Errorf("Error: %v", err)
// 	}
// 	g.AddEdge(0, 1)
// 	g.AddEdge(0, 2)
// 	g.AddEdge(0, 3)
// 	g.AddEdge(1, 3)
// 	g.AddEdge(3, 4)
// 	g.Print()
// 	n, err := g.Neighbors(0)
// 	if err != nil {
// 		t.Errorf("Error: %v", err)
// 	}
// 	fmt.Println(n)
// }

// func TestGraphBFS(t *testing.T) {
// 	g, err := NewGraph(5)
// 	if err != nil {
// 		t.Errorf("Error: %v", err)
// 	}
// 	g.AddEdge(0, 1)
// 	g.AddEdge(0, 2)
// 	g.AddEdge(0, 3)
// 	g.AddEdge(1, 3)
// 	g.AddEdge(3, 4)
// 	g.Print()
// 	found, path, err := g.BFS(0, 4)
// 	if err != nil {
// 		t.Errorf("Error: %v", err)
// 	}
// 	if found {
// 		fmt.Println(path)
// 	} else {
// 		fmt.Println("Vertex not found")
// 	}
// }

func TestGraphDFS(t *testing.T) {
	g, err := NewGraph(5)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(0, 3)
	g.AddEdge(1, 3)
	g.AddEdge(3, 4)
	g.Print()
	found, path, err := g.DFS(0, 4)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	if found {
		fmt.Println(path)
	} else {
		fmt.Println("Vertex not found")
	}
}
