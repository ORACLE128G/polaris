package main

import (
	"fmt"
	"polaris/retriever/mooc"
)

type Retriever interface {
	Get(v string) string
}

func download(r Retriever) string {
	return r.Get("https://github.com")
}

func inspect(r Retriever) {
	switch r.(type) {
	case mooc.HttpRetriever:
		fmt.Println("Type is mooc.HttpRetriever.")
	case mooc.Retriever:
		fmt.Println("Type is mooc.Retriever.")
	default:
		panic("Error occurred.")
	}
}

func typeAssertion(r Retriever) {
	retriever := r.(mooc.HttpRetriever)
	fmt.Println("Type assertion:", retriever)
}
func main() {
	r := mooc.HttpRetriever{}
	// inspect(r)
	typeAssertion(r)
	fmt.Printf("%T %v", r, r)
	//fmt.Println(download(r))
}
