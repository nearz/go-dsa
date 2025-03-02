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
	res := walk(maze, wall, start, end, seen)
	fmt.Println(res)
	path := []Point{}
	for k := range seen {
		path = append(path, k)
	}
	printPath(len(maze[0]), len(maze), seen)
	return path
}

func walk(maze []string, wall string, cur, end Point, seen map[Point]struct{}) bool {
	// off the map
	if (cur.x < 0 || cur.x >= len(maze[0])) || (cur.y < 0 || cur.y >= len(maze)) {
		return false
	}
	// a wall
	if string(maze[cur.y][cur.x]) == "#" {
		return false
	}
	// is the end
	if cur == end {
		return true
	}
	// seen
	if _, ok := seen[cur]; ok {
		return false
	}

	seen[cur] = struct{}{}
	var res bool
	// for each direction call walk with new cur as next direction
	for _, v := range dirs {
		next := Point{
			x: cur.x + v[0],
			y: cur.y + v[1],
		}
		res = walk(maze, wall, next, end, seen)
	}
	return res
}

func printPath(w, h int, seen map[Point]struct{}) {
	for ih := range h {
		for iw := range w {
			p := Point{x: iw, y: ih}
			if _, ok := seen[p]; ok {
				fmt.Print(".")
			} else {
				fmt.Print("#")
			}

		}
		fmt.Println("")
	}
}
