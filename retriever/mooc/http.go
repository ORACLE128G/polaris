package mooc

import (
	"net/http"
	"net/http/httputil"
	"time"
)

type HttpRetriever struct {
	UserAgent string
	Timeout   time.Duration
}

func (r HttpRetriever) Get(v string) string {
	resp, err := http.Get(v)
	if err != nil {
		panic("HTTP request failed.")
	}
	res, err := httputil.DumpResponse(resp, true)
	resp.Body.Close()
	if err != nil {
		panic(err)
	}
	return string(res)
}
