package graphal

import (
	"container/heap"
	"errors"
	"fmt"
	"math"
	"slices"
)

type Edge struct {
	to     int
	weight int
}

type WDGraph struct {
	list [][]Edge
	size int
}

func NewWDGraph(size int) (*WDGraph, error) {
	if size <= 0 {
		return nil, errors.New("graph size must be > 0")
	}
	list := make([][]Edge, size)
	return &WDGraph{
		list: list,
		size: size,
	}, nil
}

func (g *WDGraph) AddEdge(v1, v2, w int) error {
	if !g.inBounds(v1) {
		return errors.New("v1 out of bounds")
	}
	if !g.inBounds(v2) {
		return errors.New("v2 out of bounds")
	}
	if w < 1 {
		return errors.New("weight must be > 0")
	}
	e := Edge{to: v2, weight: w}
	g.list[v1] = append(g.list[v1], e)
	return nil
}

func (g *WDGraph) HasEdge(v1, v2 int) (bool, int) {
	if g.inBounds(v1) {
		for _, e := range g.list[v1] {
			if e.to == v2 {
				return true, e.weight
			}
		}
	}
	return false, -1
}

func (g *WDGraph) RemoveEdge(v1, v2 int) bool {
	if g.inBounds(v1) {
		for i, e := range g.list[v1] {
			if e.to == v2 {
				g.list[v1][i] = g.list[v1][len(g.list[v1])-1]
				g.list[v1] = g.list[v1][:len(g.list[v1])-1]
				return true
			}
		}
	}
	return false
}

func (g *WDGraph) Neighbors(v int) ([]Edge, error) {
	if !g.inBounds(v) {
		return nil, errors.New("vertex out of bounds")
	}
	return g.list[v], nil
}

// Dijkstra Shortest Path
func (g *WDGraph) DSP(source, target int) ([]int, error) {
	if !g.inBounds(source) {
		return nil, errors.New("source out of bounds")
	}
	if !g.inBounds(target) {
		return nil, errors.New("target out of bounds")
	}
	if source == target {
		return []int{source}, nil
	}

	prev := make([]int, g.size)
	for i := range prev {
		prev[i] = -1
	}

	dst := make([]int, g.size)
	for i := range dst {
		dst[i] = math.MaxInt
	}
	dst[source] = 0

	pq := &PriorityQ{}
	heap.Init(pq)
	heap.Push(pq, &Item{v: source, distance: 0})

	for pq.Len() > 0 {
		item := heap.Pop(pq).(*Item)
		if item.distance != dst[item.v] {
			continue
		}
		if item.v == target {
			break
		}
		n, _ := g.Neighbors(item.v)
		for _, e := range n {
			tempDst := item.distance + e.weight
			if tempDst < dst[e.to] {
				prev[e.to] = item.v
				dst[e.to] = tempDst
				heap.Push(pq, &Item{v: e.to, distance: tempDst})
			}
		}
	}

	if prev[target] == -1 {
		return nil, nil
	}

	out := []int{}
	curr := target
	for curr != -1 {
		out = append(out, curr)
		curr = prev[curr]
	}
	slices.Reverse(out)
	return out, nil
}

// Priority Queue
type Item struct {
	v        int
	distance int
}

type PriorityQ []*Item

func (pq PriorityQ) Len() int { return len(pq) }

func (pq PriorityQ) Less(i, j int) bool {
	return pq[i].distance < pq[j].distance
}

func (pq PriorityQ) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQ) Push(x any) {
	*pq = append(*pq, x.(*Item))
}

func (pq *PriorityQ) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

// helper funcs

func (g *WDGraph) inBounds(v int) bool {
	if 0 <= v && v < g.size {
		return true
	}
	return false
}

// Uility Functions
func (g *WDGraph) Print() {
	for i := range g.list {
		fmt.Printf("%d: %v\n", i, g.list[i])
	}
}
