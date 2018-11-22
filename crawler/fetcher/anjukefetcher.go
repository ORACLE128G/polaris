package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var rateLimiterOfAnjuke = time.Tick(3000 * time.Millisecond)
// Fetch data from url
func AnjukeFetch(url string) ([] byte, error) {
	<-rateLimiterOfAnjuke

	// Add request header
	r, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}
	// r.Close = true
	r.Header.Add("User-Agent",
		`Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.102 Safari/537.36`)

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
