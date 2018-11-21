package parser

import (
	"polaris/crawler/fetcher"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents, err := fetcher.AnjukeFetch("https://www.anjuke.com/sy-city.html")
	if err != nil {
		panic(err)
	}
	t.Logf("contents: %v \n", string(contents))
	result := ParseCityList(contents, "")

	for _, e := range result.Requests {
		t.Logf("url: %s \n", e)
	}
}