package parser

import (
	"polaris/crawler/fetcher"
	"testing"
)

func TestParseCity(t *testing.T) {
	contents, err := fetcher.AnjukeFetch("https://as.fang.anjuke.com/")
	if err != nil {
		panic(err)
	}
	result := ParseCity(contents, "")

	for _, e := range result.Requests {
		t.Logf("url: %s", e.Url)
	}

	t.Logf("expected url size: %d, actual: %d \n", 30, len(result.Requests))

}