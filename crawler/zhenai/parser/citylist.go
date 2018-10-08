package parser

import (
	"polaris/crawler/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents,-1)

	result := engine.ParseResult{}
	limit := 0
	for _,m := range matches{
		if limit == 10 {
			break
		}
		limit++
		result.Items = append(result.Items,"City "+ string(m[2]))
		result.Requests = append(result.Requests,engine.Request{
			Url : string(m[1]),
			ParserFunc:ParseCity,
		})
	}
	return result
}