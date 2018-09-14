package main

import (
	"fmt"
	"polaris/queue"
)

func main() {
	linkeds := queue.Linked{0, 1, 2,}
	empty := linkeds.IsEmpty()
	fmt.Println("Linked length is :?", empty)
	linkeds.Push(3)
	fmt.Println("Push a ele:", linkeds)
	linkeds.Pop()
	fmt.Println("Popping head ele:", linkeds)

}
