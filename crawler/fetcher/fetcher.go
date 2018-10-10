package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/http"
	"time"
)

var rateLimiter = time.Tick(15 * time.Millisecond)
// Fetch data from url
func Fetch(url string) ([] byte, error) {
	<-rateLimiter
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	bodyContents := resp.Body
	defer bodyContents.Close()
	if resp.StatusCode != http.StatusOK {

		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
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
