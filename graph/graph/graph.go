package graph

import (
	"container/heap"
	"errors"
	"fmt"
	"slices"

	"github.com/nearz/go-dsa/queues/queue"
)

type Vertex[T any] struct {
	// idx int
	obj T
}

func NewVertex[T any](obj T) *Vertex[T] {
	return &Vertex[T]{obj: obj}
}

type edge[T any] struct {
	to     *Vertex[T]
	weight int
}

// type weight interface{ ~int | ~float64 }
// type weight int

// TODO: Could adjList be a nested map to improve performance?
type Graph[T any] struct {
	adjList    map[*Vertex[T]][]edge[T]
	size       int
	negWeights bool
}

func New[T any]() *Graph[T] {
	return &Graph[T]{
		adjList:    make(map[*Vertex[T]][]edge[T]),
		size:       0,
		negWeights: false,
	}
}

func (g *Graph[T]) Size() int {
	return g.size
}

func (g *Graph[T]) AddEdge(v1, v2 *Vertex[T], w int) error {
	if !g.negWeights && w < 0 {
		return errors.New("graph instance does not support negative weights")
	}

	e := edge[T]{
		to:     v2,
		weight: w,
	}

	_, ok1 := g.adjList[v1]
	_, ok2 := g.adjList[v2]

	if !ok1 && !ok2 {
		g.adjList[v1] = append(g.adjList[v1], e)
		g.adjList[v2] = nil
		g.size += 2
		return nil
	}

	if ok1 && !ok2 {
		g.adjList[v1] = append(g.adjList[v1], e)
		g.adjList[v2] = nil
		g.size++
		return nil
	}

	if !ok1 && ok2 {
		g.adjList[v1] = append(g.adjList[v1], e)
		g.size++
		return nil
	}

	if ok1 && ok2 {
		g.adjList[v1] = append(g.adjList[v1], e)
		return nil
	}
	return errors.New("something went wrong")
}

func (g *Graph[T]) HasEdge(v1, v2 *Vertex[T]) bool {
	if _, ok := g.adjList[v1]; !ok {
		return false
	} else {
		for _, e := range g.adjList[v1] {
			if e.to == v2 {
				return true
			}
		}
	}
	return false
}

func (g *Graph[T]) RemoveEdge(v1, v2 *Vertex[T]) bool {
	if !g.HasEdge(v1, v2) {
		return false
	}
	for i, e := range g.adjList[v1] {
		if e.to == v2 {
			g.adjList[v1][i] = g.adjList[v1][len(g.adjList[v1])-1]
			g.adjList[v1] = g.adjList[v1][:len(g.adjList[v1])-1]
			return true
		}
	}
	return false
}

func (g *Graph[T]) Neighbors(v *Vertex[T]) []edge[T] {
	if _, ok := g.adjList[v]; ok {
		return g.adjList[v]
	}
	return []edge[T]{}
}

func (g *Graph[T]) BFS(source, target *Vertex[T]) (bool, []*Vertex[T], error) {
	if source == target {
		return true, []*Vertex[T]{source}, nil
	}

	seen := make(map[*Vertex[T]]struct{})
	seen[source] = struct{}{}

	prev := make(map[*Vertex[T]]*Vertex[T])

	q := queue.New[*Vertex[T]]()
	q.Enqueue(source)

	for !q.Empty() {
		curr, err := q.Deque()
		if err != nil {
			return false, nil, err
		}
		if curr == target {
			break
		}
		n := g.Neighbors(curr)
		for _, v := range n {
			if _, ok := seen[v.to]; ok {
				continue
			}
			seen[v.to] = struct{}{}
			prev[v.to] = curr
			q.Enqueue(v.to)
		}
	}

	if _, ok := seen[target]; !ok {
		return false, nil, nil
	}

	path := make([]*Vertex[T], 0)
	for v := target; v != nil; v = prev[v] {
		path = append(path, v)
		if v == source {
			break
		}
	}

	if path[len(path)-1] != source {
		return false, nil, nil
	}

	slices.Reverse(path)
	return true, path, nil
}

func (g *Graph[T]) DSP(source, target *Vertex[T]) (bool, []*Vertex[T]) {
	dst := make(map[*Vertex[T]]int)
	dst[source] = 0

	prev := make(map[*Vertex[T]]*Vertex[T])

	pq := &PriorityQ[T]{}
	heap.Push(pq, &Item[T]{v: source, distance: 0})

	for pq.Len() > 0 {
		curr := heap.Pop(pq).(*Item[T])
		_, ok := dst[curr.v]
		if ok && curr.distance != dst[curr.v] {
			continue
		}
		if curr.v == target {
			break
		}
		n := g.Neighbors(curr.v)
		for _, e := range n {
			tempDist := dst[curr.v] + e.weight
			_, ok := dst[e.to]
			if !ok {
				prev[e.to] = curr.v
				dst[e.to] = tempDist
				heap.Push(pq, &Item[T]{v: e.to, distance: tempDist})
				continue
			}
			if tempDist < dst[e.to] {
				prev[e.to] = curr.v
				dst[e.to] = tempDist
				heap.Push(pq, &Item[T]{v: e.to, distance: tempDist})
			}
		}

	}

	if _, ok := dst[target]; !ok {
		return false, nil
	}

	path := make([]*Vertex[T], 0)
	for v := target; v != nil; v = prev[v] {
		path = append(path, v)
		if v == source {
			break
		}
	}

	if path[len(path)-1] != source {
		return false, nil
	}

	slices.Reverse(path)
	return true, path
}

// Priority Queue
type Item[T any] struct {
	v        *Vertex[T]
	distance int
}

type PriorityQ[T any] []*Item[T]

func (pq PriorityQ[T]) Len() int { return len(pq) }

func (pq PriorityQ[T]) Less(i, j int) bool {
	return pq[i].distance < pq[j].distance
}

func (pq PriorityQ[T]) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQ[T]) Push(x any) {
	*pq = append(*pq, x.(*Item[T]))
}

func (pq *PriorityQ[T]) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

// Utility Funcs
func (g *Graph[T]) PrintVerticies() {
	for k, v := range g.adjList {
		fmt.Printf("%v: ", k)
		for i, e := range v {
			if i == 0 && i == len(v)-1 {
				fmt.Printf("[%v]", e.to.obj)
			} else if i == len(v)-1 {
				fmt.Printf(" %v]", e.to.obj)
			} else if i == 0 {
				fmt.Printf("[%v,", e.to.obj)
			} else {
				fmt.Printf(" %v,", e.to.obj)
			}
		}
		fmt.Println("")
	}
}

func PrintEdges[T any](edgs []edge[T]) {
	for i, e := range edgs {
		if i == 0 && i == len(edgs)-1 {
			fmt.Printf("[%v]", e.to.obj)
		} else if i == len(edgs)-1 {
			fmt.Printf(" %v]", e.to.obj)
		} else if i == 0 {
			fmt.Printf("[%v,", e.to.obj)
		} else {
			fmt.Printf(" %v,", e.to.obj)
		}
	}
	fmt.Println("")
}

func PrintVertexSlice[T any](vrts []*Vertex[T]) {
	for i, v := range vrts {
		if i == 0 && i == len(vrts)-1 {
			fmt.Printf("[%v]", v.obj)
		} else if i == len(vrts)-1 {
			fmt.Printf(" %v]", v.obj)
		} else if i == 0 {
			fmt.Printf("[%v,", v.obj)
		} else {
			fmt.Printf(" %v,", v.obj)
		}
	}
	fmt.Println("")
}
