package main

import (
	"fmt"
	"polaris/retriever/mooc"
)

type Retriever interface {
	Get(v string) string
}

type PostRetriever interface {
	Post(v map[string]string) string
}

type MixRetriever interface {
	PostRetriever
	Retriever
}
func download(r Retriever) string {
	return r.Get("https://github.com")
}

func post(r MixRetriever) string{
	return r.Post(map[string]string{
		"data": "Real Retriever.",
	})
}
func inspect(r Retriever) {
	// Type switch.
	switch r.(type) {
	case *mooc.HttpRetriever:
		fmt.Println("Type is mooc.HttpRetriever.")
	case mooc.Retriever:
		fmt.Println("Type is mooc.Retriever.")
	default:
		panic("Error occurred.")
	}
}
// Type Assert.
func typeAssertion(r Retriever) {
	retriever := r.(*mooc.HttpRetriever)
	fmt.Println("Type assertion:", retriever)
}
func main() {

	r := mooc.HttpRetriever{}
	// Interface comb.
	fmt.Println(post(&r))
	// inspect(r)
	//typeAssertion(&r)
	//fmt.Printf("%T %v", r, r)
	//fmt.Println(download(r))
}
