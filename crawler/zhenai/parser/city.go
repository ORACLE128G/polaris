package parser

import (
	"polaris/crawler/engine"
	"regexp"
)

var profileRe = regexp.MustCompile(
	`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a></th>`)
var cityUrlRe = regexp.MustCompile(
	`"(http://www.zhenai.com/zhenghun/[^"]+)"`)

func ParseCity(contents []byte) engine.ParseResult {
	match := profileRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range match {
		url := string(m[1])
		name := string(m[2])
		result.Requests = append(result.Requests, engine.Request{
			Url:        url,
			ParserFunc: func(v []byte) engine.ParseResult {
				return ParseProfile(v, name)
			},
		})
	}
	match = cityUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range match {
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseCity,
		})
	}

	return result
}
