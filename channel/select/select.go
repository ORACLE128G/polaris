package main

import "fmt"

func main() {
	var c1, c2  chan int
	select {
	case n := <-c1:
		fmt.Println("Received from c1", n)
	case n := <-c2:
		fmt.Println("Received from c2", n)
	default:
		fmt.Println("No value received")
	}
}
