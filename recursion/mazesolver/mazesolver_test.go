package mazesolver

import (
	"fmt"
	"testing"
)

func TestSolver(t *testing.T) {
	maze := []string{
		"#####E#",
		"#     #",
		"#S#####",
	}
	start := Point{x: 1, y: 2}
	end := Point{x: 0, y: 5}
	path := Solver(maze, "#", start, end)
	fmt.Println(path)
}

func TestSolverTwo(t *testing.T) {
	maze := []string{
		"#####E#",
		"#     #",
		"#  ## #",
		"#  #  #",
		"#S#####",
	}
	start := Point{x: 1, y: 2}
	end := Point{x: 0, y: 5}
	path := Solver(maze, "#", start, end)
	fmt.Println(path)
}
