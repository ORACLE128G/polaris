package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil {
		panic(err)
	}

	result := ParseCityList(contents, "")

	const resultSize = 470
	expectedUrls := [] string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}

	expectedCities := [] string{
		"阿坝", "阿克苏", "阿拉善盟",
	}

	if len(result.Requests) != resultSize {
		t.Errorf("result not match, should hava %d requests, but had %d",
			resultSize, len(result.Requests))
	}

	for i, url := range expectedUrls {
		if url != result.Requests[i].Url {
			t.Errorf("expected url %s, but was %s", url, result.Requests[i].Url)
		}
	}
	if len(result.Items) != resultSize {
		t.Errorf("result not match, should hava %d items, but had %d",
			resultSize, len(result.Items))
	}

	/*for i, city := range expectedCities {
		if city != result.Items[i].(string) {
			t.Errorf("expected city %s, but was %s", city, result.Items[i])
		}
	}*/
	// fmt.Printf("%s \n", contents)
}
