package _func

import "fmt"

type FuncNodes struct {
	Value string
	Left  *FuncNodes
	Right *FuncNodes
}

func (f *FuncNodes) initial(v string, l *FuncNodes, r *FuncNodes) {
	f.Value = v
	f.Left = l
	f.Right = r
}

func AutoInit() *FuncNodes{
	f := &FuncNodes{"1", nil, nil}
	f.Left = &FuncNodes{"2", nil, nil}
	f.Right = &FuncNodes{"3", nil, nil}
	return f
}

func (f *FuncNodes)Print() {
	fmt.Println(f.Value)
}

func (f *FuncNodes) EachNodes(fun func(nodes *FuncNodes)) {

	if f == nil {
		return
	}

	f.Left.EachNodes(fun)
	fun(f)
	f.Right.EachNodes(fun)
}


func Travers(nodes *FuncNodes) {
	nodes.EachNodes(func(nodes *FuncNodes) {
		nodes.Print()
	})
}
