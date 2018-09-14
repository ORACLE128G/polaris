package main

import (
	"fmt"
	"polaris/retriever/mooc"
)

type Retriever interface {
	Get(v string) string
}

func download (r Retriever) string {
	return r.Get("https://github.com")
}
func main() {
	fmt.Println(download(mooc.HttpRetriever{}))
}
