package main

import (
	"net/http"
	"polaris/fileserver"
)

func main() {
	http.HandleFunc(fileserver.Path,
		fileserver.ErrorWrapper(fileserver.HandleFuncList))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}

}
