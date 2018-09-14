package main

import (
	"fmt"
	"polaris/entry"
)

func main() {

	java := entry.Java{}
	fmt.Println(java.Version)

	node := entry.ExtTreeNode{}
	node.Pos(1)
}
