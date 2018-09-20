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

	var channels  [10]chan int
	for i := 0; i < 10; i++ {
		channels[i] = make(chan int)
		go worker(i, channels[i])
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' +i
	}
	time.Sleep(time.Minute)
}
