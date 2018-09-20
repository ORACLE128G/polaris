package main

import (
	"fmt"
	"time"
)

func TryChan() {
	c := make(chan int)
	go func() {
		for {
			n := <-c
			fmt.Println(n)
		}
	}()
	c <- 1
	c <- 2
	c <- 3
	time.Sleep(time.Microsecond)
}

func worker(id int, ch chan int) {
	for {
		fmt.Printf("Worker %d received %d \n", id, <-ch)
	}
}
func main() {
	// TryChan()
	ch := make(chan int)

	go worker(1, ch)
	ch <- 1
	ch <- 2
	time.Sleep(time.Minute)
}
