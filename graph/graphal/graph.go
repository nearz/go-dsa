package graphal

import (
	"errors"
	"fmt"
	"slices"

	"github.com/nearz/go-dsa/queues/queue"
)

type Graph struct {
	list [][]int
	size int
}

func NewGraph(size int) (*Graph, error) {
	if size <= 0 {
		return nil, errors.New("graph size must be > 0")
	}
	list := make([][]int, size)
	return &Graph{
		list: list,
		size: size,
	}, nil
}

func (g *Graph) AddEdge(v1, v2 int) error {
	if !g.inBounds(v1) {
		return errors.New("v1 out of bounds")
	}
	if !g.inBounds(v2) {
		return errors.New("v2 out of bounds")
	}
	if v1 == v2 {
		return errors.New("no self loops")
	}
	g.list[v1] = append(g.list[v1], v2)
	g.list[v2] = append(g.list[v2], v1)
	return nil
}

func (g *Graph) HasEdge(v1, v2 int) bool {
	if g.inBounds(v1) && g.inBounds(v2) {
		if v1 != v2 && slices.Contains(g.list[v1], v2) {
			return true
		}
	}
	return false
}

func (g *Graph) RemoveEdge(v1, v2 int) bool {
	if g.inBounds(v1) && g.inBounds(v2) {
		if v1 != v2 {
			if slices.Contains(g.list[v1], v2) {
				g.removeEdge(v1, v2)
				g.removeEdge(v2, v1)
				return true
			}
		}
	}
	return false
}

func (g *Graph) Neighbors(v int) ([]int, error) {
	if !g.inBounds(v) {
		return nil, errors.New("vertex out of bounds")
	}
	return g.list[v], nil
}

func (g *Graph) BFS(source, target int) (bool, []int, error) {
	if !g.inBounds(source) {
		return false, nil, errors.New("source out of bounds")
	}
	if !g.inBounds(target) {
		return false, nil, errors.New("target out of bounds")
	}
	if source == target {
		out := []int{source}
		return true, out, nil
	}

	seen := make([]bool, g.size)
	seen[source] = true
	prev := make([]int, g.size)
	for i := range prev {
		prev[i] = -1
	}
	q := queue.New[int]()
	q.Enqueue(source)

	for !q.Empty() {
		curr, err := q.Deque()
		if err != nil {
			return false, nil, err
		}
		if curr == target {
			break
		}
		for _, v := range g.list[curr] {
			if seen[v] {
				continue
			}
			seen[v] = true
			prev[v] = curr
			q.Enqueue(v)
		}
	}
	curr := target
	if prev[curr] == -1 {
		return false, nil, nil
	}
	out := []int{}
	for prev[curr] != -1 {
		out = append(out, curr)
		curr = prev[curr]
	}
	out = append(out, source)
	slices.Reverse(out)
	return true, out, nil
}

func (g *Graph) DFS(source, target int) (bool, []int, error) {
	if !g.inBounds(source) {
		return false, nil, errors.New("source not in bounds")
	}
	if !g.inBounds(target) {
		return false, nil, errors.New("target not in bounds")
	}
	if source == target {
		return true, []int{source}, nil
	}

	seen := make([]bool, g.size)
	path := []int{}

	if !g.dfs(source, target, &path, seen) {
		return false, nil, nil
	}

	return true, path, nil
}

func (g *Graph) dfs(curr, target int, path *[]int, seen []bool) bool {
	if seen[curr] {
		return false
	}
	seen[curr] = true

	*path = append(*path, curr)
	if curr == target {
		return true
	}

	for _, v := range g.list[curr] {
		if g.dfs(v, target, path, seen) {
			return true
		}
	}

	*path = (*path)[:len(*path)-1]
	return false
}

// Helper functions
func (g *Graph) removeEdge(v1, v2 int) {
	for i, e := range g.list[v1] {
		if e == v2 {
			g.list[v1][i] = g.list[v1][len(g.list[v1])-1]
			g.list[v1] = g.list[v1][:len(g.list[v1])-1]
			return

		}
	}
}

func (g *Graph) inBounds(v int) bool {
	if 0 <= v && v < g.size {
		return true
	}
	return false
}

// Utility Functions
func (g *Graph) Print() {
	for i := range g.list {
		fmt.Println(g.list[i])
	}
}
