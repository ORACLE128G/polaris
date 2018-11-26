package parser

import (
	"log"
	"polaris/crawler/fetcher"
	"testing"
)

const seed = "https://hhht.fang.anjuke.com/loupan/443998.html?from=AF_RANK_2"

func TestNewProfileParser(t *testing.T) {
	bytes, err := fetcher.AnjukeFetch(seed)
	if err != nil {
		panic(err)
	}
	parseResult := NewProfileParser().Parse(&bytes, seed)
	for r := range parseResult.Items {
		log.Printf("got item %v \n ", r)
	}
}
