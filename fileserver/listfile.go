package fileserver

import (
	"io/ioutil"
	"net/http"
	"os"
)

const (
	Path = "/list/"
)

type appHandle func(writer http.ResponseWriter, request *http.Request) error

func ErrorWrapper(handle appHandle) func(writer http.ResponseWriter,
	request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		err := handle(writer, request)
		if err != nil {
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}

func HandleFuncList(writer http.ResponseWriter, request *http.Request) error {
	s := request.URL.Path[len(Path):]
	file, e := os.Open(s)
	if e != nil {
		return e
	} else {
		defer file.Close()
		bytes, i := ioutil.ReadAll(file)
		if i != nil {
			panic(i)
		} else {
			writer.Write(bytes)
		}
	}
	return nil
}
