package entry

import "polaris/entry/node"

type ExtTreeNode struct {
	tree *node.Tree
	v    int
}

func (tree *ExtTreeNode) Pos(i int) int {
	if tree == nil {
		return int(1)
	}
	nodes := ExtTreeNode{}
	return nodes.v
}
