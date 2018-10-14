package main

import (
	"net/http"
	"polaris/crawler/frontend/controller"
)

func main() {
	http.Handle("/", http.FileServer(
		http.Dir("crawler/frontend/view")))
	http.Handle("/search", controller.CreateSearchResultHandler(
		"crawler/frontend/view/template.html"))
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		panic(err)
	}
}
