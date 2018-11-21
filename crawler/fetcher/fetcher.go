package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var rateLimiter = time.Tick(80 * time.Millisecond)
// Fetch data from url
func Fetch(url string) ([] byte, error) {
	<-rateLimiter

	// Add request header
	r, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}
	r.Header.Add("User-Agent",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.102 Safari/537.36")
	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			log.Printf("%s Redirect to %s \n", req.Referer(), req.URL)
			return nil
		},
	}
	resp, err := client.Do(r)
	/*if err != nil{
		panic(err)
	}
	resp, err := http.Get(url)*/
	if err != nil {
		return nil, err
	}
	bodyContents := resp.Body
	defer bodyContents.Close()
	if resp.StatusCode != http.StatusOK {

		return nil, fmt.Errorf("wrong status code: %d \n", resp.StatusCode)
	}
	bodyReader := bufio.NewReader(bodyContents)
	e := determineCharset(bodyReader)

	return ioutil.ReadAll(transform.NewReader(bodyReader, e.NewDecoder()))
}

// Determine Charset from r
func determineCharset(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
