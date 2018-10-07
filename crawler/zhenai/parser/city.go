package parser

import (
	"polaris/crawler/engine"
	"regexp"
)

var ProfileRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a></th>`)
var cityUrlRe = regexp.MustCompile(`"(http://www.zhenai.com/zhenghun/[0-9[a-z][A-Z]/[\d]+)"`)

func ParseCity(contents []byte) engine.ParseResult {
	match := ProfileRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range match {
		url := string(m[1])
		name := string(m[2])
		result.Items = append(result.Items, "User "+name)
		result.Requests = append(result.Requests, engine.Request{
			Url:        url,
			ParserFunc: engine.NilParser,
		})
	}
	match = cityUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range match {
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: engine.NilParser,
		})
	}

	return result
}
