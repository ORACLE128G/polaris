package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
)

// Determine Charset from r
func determineCharset(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(r)
	}

	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
func main() {

	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error code: %d", resp.StatusCode)
		return
	}
	e := determineCharset(resp.Body)

	body, err := ioutil.ReadAll(transform.NewReader(resp.Body, e.NewDecoder()))
	if err != nil {
		panic(body)
	}
	fmt.Printf("%s \n ", body)
}
