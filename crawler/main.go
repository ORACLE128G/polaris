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
	"regexp"
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

// Matching all available city.
func cityList(b [] byte) {
	//reg := "(http://www.zhenai.com/zhenghun/[0-9a-z]+)([^<]+)"
	reg := `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`
	c := regexp.MustCompile(reg)
	matches := c.FindAllSubmatch(b, -1)
	for _, m := range matches {
		fmt.Printf("City: %s, URL: %s \n",m[2],m[1])
	}
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
	cityList(body)
}
