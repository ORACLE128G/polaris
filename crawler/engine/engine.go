package engine

import (
	"log"
	"polaris/crawler/fetcher"
)

func Run(seeds ...Request) {
	var requests [] Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	bk := 0
	for len(requests) > 0 {

		// Fetch 10 cities from requests.
		if bk == 10 {
			break
		}
		bk++
		// Get head requests
		r := requests[0]
		requests = requests[1:]

		log.Printf("Fetching %s", r.Url)
		body, err := fetcher.	Fetch(r.Url)
		if err != nil {
			log.Printf("Fetching: error fetching url %s: %v", r.Url, err)
			continue
		}
		parseResult	 := r.ParserFunc(body)
		requests = append(requests, parseResult.Requests...)

		for _, i := range parseResult.Items {
			log.Printf("Got item %s", i)
		}

	}
}
