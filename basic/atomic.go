package main

import (
	"fmt"
	"sync"
	"time"
)

type atomicInt struct {
	v int
	lock sync.Mutex
}

func (a *atomicInt) increment(){
	fmt.Println("Safe increment.")
	func (){
		a.lock.Lock()
		defer a.lock.Unlock()
		a.v++
	}()

}

func (a *atomicInt) get() int{
	a.lock.Lock()
	defer a.lock.Unlock()
	return int(a.v)
}
func main() {
	var a atomicInt
	a.increment()
	go func() {
		a.increment()
	}()
	time.Sleep(time.Second)
	fmt.Println(a.get())
}
