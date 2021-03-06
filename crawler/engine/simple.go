package engine

import (
	"log"
	"polaris/crawler/fetcher"
)

type SimpleEngine struct {
}

func (s SimpleEngine) Run(seeds ...Request) {
	var requests [] Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		// Get head requests
		r := requests[0]
		requests = requests[1:]

		parseResult, err := Worker(r)
		if err != nil {
			continue
		}
		requests = append(requests, parseResult.Requests...)

		for _, i := range parseResult.Items {
			log.Printf("Got item %v", i)
		}

	}
}

func Worker(r Request) (ParseResult, error) {
	log.Printf("Fetching %s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetching %s: %v \n", r.Url, err)
		return ParseResult{}, err
	}
	return r.Parser.Parse(body, r.Url), nil
}

func AnjukeWorker(r Request) (ParseResult, error){
	log.Printf("Fetching %s", r.Url)
	body, err := fetcher.AnjukeFetch(r.Url)
	if err != nil {
		log.Printf("Fetching %s: %v \n", r.Url, err)
		return ParseResult{}, err
	}
	return r.Parser.Parse(body, r.Url), nil
}
