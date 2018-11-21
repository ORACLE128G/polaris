package parser

import (
	"polaris/crawler-distributed/config"
	"polaris/crawler/engine"
	"regexp"
)

const (
	cityListRe      = `<a href="https://([0-9a-z]+).anjuke.com">([^<]+)</a>`
	basicListPrefix = `.fang.anjuke.com`
	httpsPrefix     = `https://`
)

func ParseCityList(contents []byte, _ string) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url: httpsPrefix + string(m[1]) + basicListPrefix,
			Parser: engine.NewFuncParser(
				ParseCity, config.ParseCity),
		})
	}
	return result
}
