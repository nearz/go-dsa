package graph_am

import (
	"errors"
	"fmt"
	"math/bits"
	"slices"
	"strconv"
	"strings"

	"github.com/nearz/go-dsa/queues/queue"
)

// Graph is an undirected unweighted graph implemented with a uint64 matrix
type Graph struct {
	matrix [][]uint64
	size   int
}

func NewGraph(size int) (*Graph, error) {
	if size <= 0 {
		return nil, errors.New("graph size must be > 0")
	}
	wordsPerRow := (size + 63) >> 6
	m := make([][]uint64, size)
	for i := range m {
		m[i] = make([]uint64, wordsPerRow)
	}
	return &Graph{
		matrix: m,
		size:   size,
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
	g.setBit(v1, v2)
	g.setBit(v2, v1)

	return nil
}

func (g *Graph) RemoveEdge(v1, v2 int) bool {
	if g.inBounds(v1) && g.inBounds(v2) {
		if v1 != v2 {
			if !g.bitExists(v1, v2) {
				return false
			}
			g.clearBit(v1, v2)
			g.clearBit(v2, v1)

			return true
		}
	}
	return false
}

func (g *Graph) HasEdge(v1, v2 int) bool {
	if g.inBounds(v1) && g.inBounds(v2) {
		if v1 != v2 && g.bitExists(v1, v2) {
			return true
		}
	}
	return false
}

func (g *Graph) Neighbors(source int) ([]int, error) {
	if !g.inBounds(source) {
		return nil, errors.New("vertex out of bounds")
	}
	row := g.matrix[source]
	n := []int{}
	for w, word := range row {
		x := word
		for x != 0 {
			idx := bits.TrailingZeros64(x)
			col := (w << 6) + idx
			if col < g.size {
				n = append(n, col)
			}
			x &= x - 1
		}
	}
	return n, nil
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
	prev := make([]int, g.size)
	for i := range prev {
		prev[i] = -1
	}
	q := queue.New[int]()
	q.Enqueue(source)
	seen[source] = true

	for !q.Empty() {
		curr, err := q.Deque()
		if err != nil {
			return false, nil, err
		}
		if curr == target {
			break
		}
		n, err := g.Neighbors(curr)
		if err != nil {
			return false, nil, err
		}
		for _, v := range n {
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
		return false, nil, errors.New("source out of bounds")
	}
	if !g.inBounds(target) {
		return false, nil, errors.New("target out of bounds")
	}
	if source == target {
		return true, []int{source}, nil
	}

	prev := make([]int, g.size)
	for i := range prev {
		prev[i] = -1
	}
	seen := make([]bool, g.size)

	err := g.dfs(source, target, prev, seen)
	if err != nil {
		return false, nil, err
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

func (g *Graph) dfs(curr, target int, prev []int, seen []bool) error {
	if curr == target {
		seen[target] = true
		return nil
	}
	if seen[target] {
		return nil
	}

	seen[curr] = true
	n, err := g.Neighbors(curr)
	if err != nil {
		return err
	}
	for _, v := range n {
		if seen[v] {
			continue
		}
		prev[v] = curr
		g.dfs(v, target, prev, seen)
	}
	return nil
}

// Helper Functions
func (g *Graph) inBounds(v int) bool {
	if 0 <= v && v < g.size {
		return true
	}
	return false
}

func (g *Graph) bitExists(i, j int) bool {
	word := j >> 6
	mask := uint64(1) << (j & 63)
	if (g.matrix[i][word] & mask) == 0 {
		return false
	}
	return true
}

func (g *Graph) setBit(i, j int) {
	word := j >> 6
	mask := uint64(1) << (j & 63)
	g.matrix[i][word] |= mask
}

func (g *Graph) clearBit(i, j int) {
	word := j >> 6
	mask := uint64(1) << (j & 63)
	g.matrix[i][word] &^= mask
}

// Utility Functions
func (g *Graph) Print() {
	if g == nil || g.size == 0 {
		fmt.Println("empty graph")
		return
	}

	numCols := g.size

	// Calculate padding width for row labels based on the largest index
	rowLabelWidth := len(strconv.Itoa(g.size - 1))

	// Header
	fmt.Printf("%*s | ", rowLabelWidth, "")
	for c := 0; c < numCols; c++ {
		fmt.Printf("%3d", c)
	}
	fmt.Println()
	fmt.Printf("%s-+-%s\n", strings.Repeat("-", rowLabelWidth), strings.Repeat("-", numCols*3))

	// Rows
	for r := 0; r < g.size; r++ {
		fmt.Printf("%*d | ", rowLabelWidth, r)
		for c := 0; c < numCols; c++ {
			word := c >> 6
			bit := uint64(1) << uint(c&63)
			if (g.matrix[r][word] & bit) != 0 {
				fmt.Printf("\x1b[31m%3s\x1b[0m", "1")
			} else {
				fmt.Printf("%3s", "0")
			}
		}
		fmt.Println()
	}
}
