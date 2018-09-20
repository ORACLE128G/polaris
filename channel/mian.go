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

func createWorker(id int) chan<- int {
	c := make(chan int)
	go func() {
		for {
			fmt.Printf("Worker %d received %d \n", id, <-c)
		}
	}()
	return c
}
func chanDemo() {
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}
	time.Sleep(time.Minute)
}

func makeBufferedChannel() chan int{
	return make(chan int, 3)
}
func main() {
	// TryChan()
	// chanDemo()
	channel := makeBufferedChannel()
	channel <- 3
	c := <-channel
	fmt.Println( c )


}
