package graphal

import (
	"fmt"
	"testing"
)

// func TestNewWDGraph(t *testing.T) {
// 	g, err := NewWDGraph(5)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	g.Print()
// }

// func TestWDGraphAddEdges(t *testing.T) {
// 	g, err := NewWDGraph(5)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	g.AddEdge(0, 1, 2)
// 	g.AddEdge(0, 3, 3)
// 	g.AddEdge(1, 3, 4)
// 	g.AddEdge(3, 4, 2)
// 	g.AddEdge(2, 0, 1)
// 	g.Print()
// }

// func TestWDGraphHasEdge(t *testing.T) {
// 	g, err := NewWDGraph(5)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	g.AddEdge(0, 1, 2)
// 	g.AddEdge(0, 3, 3)
// 	g.AddEdge(1, 3, 4)
// 	g.AddEdge(3, 4, 2)
// 	g.AddEdge(2, 0, 1)
// 	g.Print()
// 	res, w := g.HasEdge(0, 3)
// 	if res {
// 		fmt.Printf("Graph has edge with weight: %d", w)
// 	} else {
// 		fmt.Println("Edge not found")
// 	}
// }

// func TestWDGraphNeighbors(t *testing.T) {
// 	g, err := NewWDGraph(5)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	g.AddEdge(0, 1, 2)
// 	g.AddEdge(0, 3, 3)
// 	g.AddEdge(1, 3, 4)
// 	g.AddEdge(3, 4, 2)
// 	g.AddEdge(2, 0, 1)
// 	n, err := g.Neighbors(0)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println(n)
// }

// func TestWDGraphRemoveEdge(t *testing.T) {
// 	g, err := NewWDGraph(5)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	g.AddEdge(0, 1, 2)
// 	g.AddEdge(0, 3, 3)
// 	g.AddEdge(1, 3, 4)
// 	g.AddEdge(3, 4, 2)
// 	g.AddEdge(2, 0, 1)
// 	g.Print()
// 	fmt.Println("")
// 	res := g.RemoveEdge(0, 3)
// 	if res {
// 		fmt.Println("vertex removed")
// 		g.Print()
// 	} else {
// 		fmt.Println("vertex not removed")
// 	}
// }

// func TestLowU(t *testing.T) {
// 	seen := []bool{true, false, false, false, false}
// 	dst := []int{1, 2, 3, 1, -1}
// 	i := lowU(seen, dst)
// 	fmt.Println(i)
// }

func TestDSP(t *testing.T) {
	g, err := NewWDGraph(5)
	if err != nil {
		fmt.Println(err)
		return
	}
	g.AddEdge(0, 1, 1)
	g.AddEdge(0, 2, 5)
	g.AddEdge(1, 2, 7)
	g.AddEdge(1, 3, 3)
	g.AddEdge(3, 2, 2)
	g.AddEdge(2, 4, 1)
	g.Print()
	path, err := g.DSP(0, 4)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(path)
}

func TestDSPTw0(t *testing.T) {
	g, err := NewWDGraph(5)
	if err != nil {
		fmt.Println(err)
		return
	}
	g.AddEdge(0, 1, 1)
	g.AddEdge(0, 2, 5)
	g.AddEdge(1, 2, 1)
	g.AddEdge(1, 3, 3)
	g.AddEdge(3, 2, 2)
	g.AddEdge(2, 4, 1)
	g.Print()
	path, err := g.DSP(0, 4)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(path)
}
