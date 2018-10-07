package parser

import (
	"polaris/crawler/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents [] byte) engine.ParseResult {
	//reg := "(http://www.zhenai.com/zhenghun/[0-9a-z]+)([^<]+)"
	reg := cityListRe
	c := regexp.MustCompile(reg)
	matches := c.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: engine.NilParser,
		})
	}
	return result
}
