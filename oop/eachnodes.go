package main

import (
	"fmt"
)

type Nodes struct {
	V           int
	Left, Right *Nodes
}

func (nodes *Nodes) print() {
	fmt.Println("Each node is :", nodes.V)
}
func inst() *Nodes {
	root := Nodes{0, nil, nil}
	root.Left = &Nodes{1, nil, nil}
	root.Right = &Nodes{2, nil, nil}
	root.Left.Left = &Nodes{3, nil, nil}
	root.Left.Right = &Nodes{4, nil, nil}

	return &root
}

func (nodes *Nodes) setVal(v int) {
	nodes.Left.V = v
}
func (nodes *Nodes) TraverseFunc(f func(*Nodes)) {
	if nodes == nil {
		return
	}
	nodes.Left.TraverseFunc(f)
	f(nodes)
	nodes.Right.TraverseFunc(f)
}

func (nodes *Nodes) traverseWithChannel() chan *Nodes {
	out := make(chan *Nodes)
	go func() {
		nodes.TraverseFunc(func(nodes *Nodes) {
			out <- nodes
		})
		close(out)
	}()
	return out
}
func main() {
	nodes := inst()
	//nodes.print()
	//nodes.TraverseFunc()

	maxNode := 0
	c := nodes.traverseWithChannel()
	for node := range c {
		fmt.Println(node)
		if node.V > maxNode {
			maxNode = node.V
		}
	}

	fmt.Println("Max node value: ", maxNode)
}
