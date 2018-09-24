package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func httpClient(){
	r, e := http.NewRequest(http.MethodGet, "https://www.imooc.com", nil)
	if e != nil {
		panic(e)
	}
	r.Header.Add("User-Agent","Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Mobile Safari/537.36")
	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println("Redirect:", req)
			return nil
		},
	}
	res, e := client.Do(r)
	if e != nil{
		panic(e)
	}
	defer res.Body.Close()

	body, e := httputil.DumpResponse(res, true)
	if e != nil{
		panic(e)
	}
	fmt.Printf("%s \n", body)

}
func main() {
	httpClient()
}
