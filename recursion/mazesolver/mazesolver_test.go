package mazesolver

import (
	"slices"
	"testing"
)

func TestSolver(t *testing.T) {
	expRes := []Point{
		{x: 1, y: 2},
		{x: 1, y: 1},
		{x: 2, y: 1},
		{x: 3, y: 1},
		{x: 4, y: 1},
		{x: 5, y: 1},
		{x: 5, y: 0},
	}
	maze := []string{
		"#####E#",
		"#     #",
		"#S#####",
	}
	start := Point{x: 1, y: 2}
	end := Point{x: 5, y: 0}
	path := Solver(maze, "#", start, end)
	if !slices.Equal(path, expRes) {
		t.Errorf("Resulting path incorrect: \nexpected: %v, \nresult: %v", expRes, path)
	}
}

func TestSolverTwo(t *testing.T) {
	expRes := []Point{
		{x: 1, y: 4},
		{x: 1, y: 3},
		{x: 1, y: 2},
		{x: 1, y: 1},
		{x: 2, y: 1},
		{x: 3, y: 1},
		{x: 4, y: 1},
		{x: 5, y: 1},
		{x: 5, y: 0},
	}
	maze := []string{
		"#####E#",
		"#     #",
		"# ### #",
		"# ##  #",
		"#S#####",
	}
	start := Point{x: 1, y: 4}
	end := Point{x: 5, y: 0}
	path := Solver(maze, "#", start, end)
	if !slices.Equal(path, expRes) {
		t.Errorf("Resulting path incorrect: \nexpected: %v, \nresult: %v", expRes, path)
	}
}
