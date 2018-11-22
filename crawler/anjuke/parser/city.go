package parser

import (
	"polaris/crawler-distributed/config"
	"polaris/crawler/engine"
	"regexp"
)

var profileRe = regexp.MustCompile(
	`<a class="pic"  href="(https://[a-z]+\.fang\.anjuke\.com/loupan/[a-z0-9]+\.html)"`)
var cityUrlRe = regexp.MustCompile(
	`"<a href="https://([.]+).fang.anjuke.com/loupan/([^"])">([.]+)</a>"`)

func ParseCity(contents []byte, _ string) engine.ParseResult {
	match := profileRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range match {
		url := string(m[1])
		//name := string(m[2])
		result.Requests = append(result.Requests, engine.Request{
			Url:    url,
			//Parser: NewProfileParser(name),
		})
	}
	match = cityUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range match {
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			Parser: engine.NewFuncParser(
				ParseCity,
				config.ParseCity),
		})
	}

	return result
}
