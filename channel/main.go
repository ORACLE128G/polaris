package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Response code is %d", resp.StatusCode)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(body)
	}
	fmt.Printf("%s \n ", body)
}
