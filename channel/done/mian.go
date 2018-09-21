package main

import (
	"fmt"
)

func createWorker(id int) works {
	w := works{
		in:   make(chan int),
		done: make(chan bool),
	}
	go func() {
		for {
			fmt.Printf("Worker %d received %d \n", id, <-w.in)
			go func() {
				w.done <- true
			}()
		}
	}()
	return w
}
func chanDemo() {
	var w [10] works
	for i := range w {
		w[i] = createWorker(i)
	}

	for i := range w {
		w[i].in <- 'a' + i
	}
	for i := range w {
		w[i].in <- 'A' + i
	}
	// wait for all of them.
	for _, k := range w {
		<-k.done
		<-k.done
	}
}

type works struct {
	in   chan int
	done chan bool
}

func main() {
	// TryChan()
	chanDemo()
	/*channel := makeBufferedChannel()
	channel <- 3
	c := <-channel
	fmt.Println( c )*/

}
