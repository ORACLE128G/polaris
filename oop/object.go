package main

import "fmt"

type treeNode struct {
	value       int
	left, right *treeNode
}

func instance(value int) *treeNode {

	return &treeNode{value, nil, nil}
}
func main() {
	var root treeNode
	root = treeNode{value: 3}
	root.left = &treeNode{2, nil, nil}
	root.right = &treeNode{1, nil, nil}
	root.right = new(treeNode)

	nodes := []treeNode{
		{value: 1},
		{},
		{3, nil, &root},
	}

	fmt.Println(nodes)

	node := instance(2)
	node = &treeNode{0, nil, nil}
	fmt.Println(node)

}
