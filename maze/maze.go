package main

import (
	"fmt"
	"os"
)

func readMaze() [][]int {
	file, e := os.Open("maze/maze.in")
	if e != nil {
		panic(e)
	}
	defer file.Close()
	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)
	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([] int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}
	return maze
}
func main() {
	maze := readMaze()
	for _, row := range maze {
		for _, col := range row {
			fmt.Printf("%d ",col)
		}
		fmt.Println()
	}
}
