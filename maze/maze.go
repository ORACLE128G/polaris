package main

import (
	"fmt"
	"os"
)

func readMaze(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)

	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}

	return maze
}

type point struct {
	// i is row
	// and j is col.
	i, j int
}

var dirs = [4]point{
	{-1, 0},
	{0, -1},
	{1, 0},
	{0, 1},
}

// addition point.
func (p point) add(r point) point {
	return point{p.i + r.i, p.j + r.j}
}

//
func (p point) at(grid [][]int) (int, bool) {
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}

	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}

	return grid[p.i][p.j], true
}

func home(maze [][] int, start, end point) [][]int {
	q := []point{start}
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}

	for len(q) > 0 {
		// get queue head.
		cur := q[0]
		q = q[1:]

		if cur == end {
			break
		}
		for _, dir := range dirs {
			// addition dir with 1 step.
			next := cur.add(dir)

			v, ok := next.at(maze)
			if !ok || v == 1 {
				continue
			}

			v, ok = next.at(steps)
			if !ok || v != 0 {
				continue
			}

			if next == start {
				continue
			}

			s, _ := cur.at(steps)
			steps[next.i][next.j] = s + 1
			q = append(q, next)

		}
	}
	return steps
}

// back to start.
func back(maze [][]int, start, end point) []int {
	if end == start {
		return nil
	}
	e := maze[end.i][end.j]

	cur := end
	r := []int{e}
	next := e - 1
	for maze[start.i][start.j] != next {
		for _, dir := range dirs {
			n := cur.add(dir)
			v, ok := n.at(maze)
			if !ok || v != next {
				continue
			}
			r = append(r, next)
			next -= 1
			cur = n
			break
		}
	}

	return r
}
func main() {
	maze := readMaze("maze/maze.in")

	steps := home(maze, point{0, 0}, point{len(maze) - 1, len(maze[0]) - 1})
	for _, i := range steps {
		for _, v := range i {
			fmt.Printf("%3d ", v)
		}
		fmt.Println()
	}

	r := back(steps, point{0, 0}, point{len(maze) - 1, len(maze[0]) - 1})
	fmt.Println(r)
	// TODO: Print walk maze map.
}
