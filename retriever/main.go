package main

import "polaris/retriever/mooc"

type Retriever interface {
	Get(v string) string
}

func download (r Retriever) string {
	return r.Get("mooc.com.cn")
}
func main() {
	download(mooc.Retriever{Contents:"Contents"})
}
