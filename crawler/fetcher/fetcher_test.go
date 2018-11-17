package fetcher_test

import (
	"fmt"
	"polaris/crawler/fetcher"
	"testing"
)

func TestFetch(t *testing.T) {
	b, err := fetcher.Fetch("http://album.zhenai.com/u/1347842007")
	if err != nil {
		panic(err)
	}
	fmt.Println()
	t.Logf("fetcher contents: %s \n", string(b))
	fmt.Println()
}
