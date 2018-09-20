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
func main() {
	TryChan()
}
