package mazesolver

import "fmt"

type Point struct {
	x int
	y int
}

var dirs = map[string][]int{
	"l": {-1, 0},
	"u": {0, -1},
	"r": {1, 0},
	"d": {0, 1},
}

func Solver(maze []string, wall string, start, end Point) []Point {
	seen := make(map[Point]struct{})
	path := []Point{}
	walk(maze, wall, start, end, seen, &path)
	// printPath(maze, seen)
	return path
}

func walk(maze []string, wall string, cur, end Point, seen map[Point]struct{}, path *[]Point) {
	// slice to hold next nodes for all directions.
	next := []Point{}
	for _, v := range dirs {
		nextNode := Point{
			x: cur.x + v[0],
			y: cur.y + v[1],
		}
		// if next node is end return
		// if next node is out of bounds, a wall or in seen do not add to slice
		if nextNode == end {
			seen[cur] = struct{}{}
			seen[nextNode] = struct{}{}
			*path = append(*path, cur)
			*path = append(*path, nextNode)
			return
		} else if !inBounds(nextNode, maze) || string(maze[nextNode.y][nextNode.x]) == "#" {
			continue
		} else if _, ok := seen[nextNode]; ok {
			continue
		} else {
			next = append(next, nextNode)
		}
	}

	seen[cur] = struct{}{}
	*path = append(*path, cur)
	// for all nodes in next slice call walk
	for i := range next {
		walk(maze, wall, next[i], end, seen, path)
	}
}

func inBounds(p Point, maze []string) bool {
	if (p.x < 0 || p.x >= len(maze[0])) || (p.y < 0 || p.y >= len(maze)) {
		return false
	}
	return true
}

func printPath(maze []string, seen map[Point]struct{}) {
	h := len(maze)
	w := len(maze[0])
	for y := range h {
		for x := range w {
			p := Point{x: x, y: y}
			if _, ok := seen[p]; ok {
				fmt.Print(".")
			} else {
				fmt.Print(string(maze[y][x]))
			}

		}
		fmt.Println("")
	}
}
