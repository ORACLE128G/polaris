package main

import (
	"fmt"
	"math/rand"
	"time"
)

func genChannel() chan int {

	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(
				rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func worker(id int, c chan int) {
	for n := range c {
		fmt.Printf("Received %d from %d \n", n, id)
	}
}

func createWorker(id int) chan<- int {

	c := make(chan int)
	go worker(id, c)
	return c
}
func main() {
	var c1, c2 = genChannel(), genChannel()
	w := createWorker(0)
	n := 0
	hasV := false
	for {
		var actCh chan<- int
		if hasV {
			actCh = w
		}
		select {
		case n = <-c1:
			hasV = true
		case n = <-c2:
			hasV = true
		case actCh <- n:
			hasV = false
		}
	}
}
