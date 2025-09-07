package adjmatrix

import (
	"fmt"
	"testing"
)

// func TestNew(t *testing.T) {
// 	g := New(15)
// 	g.Print()
// }

// func TestAddEdges(t *testing.T) {
// 	g := New(4)
// 	g.AddEdge(0, 1)
// 	g.AddEdge(1, 2)
// 	g.Print()
// }

func TestBFS(t *testing.T) {
	g := New(5)
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(0, 3)
	g.AddEdge(1, 3)
	g.AddEdge(3, 4)
	g.Print()
	path := g.BFS(0, 4)
	fmt.Println(path)
}

func TestBFSTwo(t *testing.T) {
	g := New(5)
	g.AddEdge(0, 1)
	g.AddEdge(0, 3)
	g.AddEdge(1, 3)
	g.AddEdge(3, 4)
	g.Print()
	path := g.BFS(0, 2)
	fmt.Println(path)
}

// func TestNeighbors(t *testing.T) {
// 	g := New(5)
// 	g.AddEdge(0, 1)
// 	g.AddEdge(0, 2)
// 	g.AddEdge(0, 3)
// 	g.AddEdge(1, 3)
// 	g.AddEdge(3, 4)
// 	n := g.Neighbors(0)
// 	fmt.Println(n)
// }

// func TestMediumDensityGraph(t *testing.T) {
// 	size := 20
// 	g := New(size)
// 	k := 3 // connect each node to its next 3 neighbors (sparser)
// 	for i := 0; i < size; i++ {
// 		for j := 1; j <= k; j++ {
// 			g.AddEdge(i, (i+j)%size)
// 		}
// 	}
//
// 	g.Print()
// 	path := g.BFS(0, 10)
// 	fmt.Println(path)
// }
