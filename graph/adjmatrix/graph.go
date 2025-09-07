package adjmatrix

import (
	"errors"
	"fmt"
	"slices"

	"github.com/nearz/go-dsa/queues/queue"
)

type Graph struct {
	matrix [][]int
	size   int
}

func New(size int) *Graph {
	m := make([][]int, size)
	for i := range m {
		m[i] = make([]int, size)
	}
	return &Graph{
		matrix: m,
		size:   size,
	}
}

func (g *Graph) AddEdge(v1, v2 int) error {
	if v1 < 0 || g.size <= v1 {
		return errors.New("V1 out of bounds")
	}
	if v2 < 0 || g.size <= v2 {
		return errors.New("V2 out of bounds")
	}
	if v1 == v2 {
		return errors.New("No self-loops")
	}
	g.matrix[v1][v2] = 1
	g.matrix[v2][v1] = 1
	return nil
}

func (g *Graph) RemoveEdge(v1, v2 int) error {
	if v1 < 0 || g.size <= v1 {
		return errors.New("V1 out of bounds")
	}
	if v2 < 0 || g.size <= v2 {
		return errors.New("V2 out of bounds")
	}
	if g.matrix[v1][v2] == 0 {
		return errors.New("No Edge")
	}
	g.matrix[v1][v2] = 0
	g.matrix[v2][v1] = 0
	return nil
}

func (g *Graph) Neighbors(source int) []int {
	if source < 0 || g.size <= source {
		return []int{}
	}
	n := []int{}
	for vrtx, v := range g.matrix[source] {
		if v != 0 {
			n = append(n, vrtx)
		}
	}
	return n
}

func (g *Graph) BFS(source, val int) []int {
	if val < 0 || g.size <= val {
		return []int{}
	}
	seen := make([]bool, g.size)
	for i := range seen {
		seen[i] = false
	}
	prev := make([]int, g.size)
	for i := range prev {
		prev[i] = -1
	}
	q := queue.New[int]()
	seen[source] = true
	q.Enqueue(source)

	for !q.Empty() {
		curr, err := q.Deque()
		if err != nil {
			return []int{}
		}
		if curr == val {
			break
		}
		for vrtx, e := range g.matrix[curr] {
			if e == 0 {
				continue
			}
			if seen[vrtx] {
				continue
			}
			seen[vrtx] = true
			prev[vrtx] = curr
			q.Enqueue(vrtx)
		}
	}

	fmt.Println(prev)
	curr := val
	if prev[curr] == -1 {
		return []int{}
	}
	out := []int{}
	for prev[curr] != -1 {
		out = append(out, curr)
		curr = prev[curr]
	}
	out = append(out, source)
	slices.Reverse(out)
	return out
}

func (g *Graph) Print() {
	if g.size == 0 {
		fmt.Println("Empty graph")
		return
	}

	// Print column indices
	fmt.Printf("    ")
	for c := 0; c < g.size; c++ {
		fmt.Printf("%3d ", c)
	}
	fmt.Println()

	// Print separator
	fmt.Printf("    ")
	for c := 0; c < g.size; c++ {
		fmt.Printf("----")
	}
	fmt.Println()

	// Print rows with row indices
	for i := 0; i < g.size; i++ {
		fmt.Printf("%3d|", i) // row index
		for j := 0; j < g.size; j++ {
			fmt.Printf("%3d ", g.matrix[i][j])
		}
		fmt.Println()
	}
}
