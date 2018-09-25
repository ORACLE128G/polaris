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


func main() {
	maze := readMaze("maze/maze.in")

	for i := range maze {
		for j := range maze[i] {
			fmt.Printf("%d ", maze[i][j])
		}
		fmt.Println()
	}


	// TODO: Print walk maze map.
}
