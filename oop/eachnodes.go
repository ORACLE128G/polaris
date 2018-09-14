package main

import "fmt"

type Nodes struct {
	v           int
	left, right *Nodes
}

func (nodes *Nodes) print () {
	fmt.Println("Each node is :",nodes.v)
}
func inst() *Nodes {
	root := Nodes{0, nil, nil}
	root.left = &Nodes{1, nil, nil}
	root.right = &Nodes{2, nil, nil}
	root.left.left = &Nodes{3, nil, nil}
	root.left.right = &Nodes{4, nil, nil}

	return &root
}

func (nodes *Nodes)setVal(v int) {
	nodes.left.v = v
}
func (nodes *Nodes)traverse()[]int {
	if nodes == nil {
		return nil
	}
	nodes.left.traverse()
	nodes.print()
	nodes.right.traverse()
	return []int{}
}
func main() {
	nodes := inst()
	nodes.print()
	nodes.traverse()
}
