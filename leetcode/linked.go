package main

import (
	"fmt"
	"time"
)

type Linked struct {
	Val  int
	Next *Linked
}

func (head *Linked) RemoveNthFromEnd(index int) {

	//c := 0
	/*go func(head *Linked) {
		fmt.Println(head)
	}(head)*/


	time.Sleep(time.Millisecond)
}

func (head *Linked) PrintAllOfThem() {

	fmt.Println(head.Val)
	if head.Next != nil {
		head.Next.PrintAllOfThem()
	}
}

func main() {
	head := &Linked{1, &Linked{2, &Linked{3, &Linked{4, nil}}}}
	//head.PrintAllOfThem()
	head.RemoveNthFromEnd(0)
}
