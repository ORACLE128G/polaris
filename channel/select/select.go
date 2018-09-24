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
				rand.Intn(3)) * time.Second)
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
	var queue  [] int
	t := time.After(10 * time.Second)
	for {
		var actCh chan<- int
		var actV int
		if len(queue) > 0 {
			actCh = w
			actV = queue[0]
		}
		select {
		case n = <-c1:
			queue = append(queue, n)
		case n = <-c2:
			queue = append(queue, n)
		case actCh <- actV:
			queue = queue[1:]
		case <- time.After(800 * time.Millisecond):
			fmt.Println("Timeout.")
		case <- t:
			fmt.Println("Bye.")
			return
		}
	}
}
