package mooc

import "fmt"

type Retriever struct {
	Contents string

}

func (r Retriever) Get (v string) string {
	fmt.Println("Retriever contents is :", r.Contents)
	return r.Contents
}