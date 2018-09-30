package main

import (
	"polaris/crawler/engine"
	"polaris/crawler/zhenai/parser"
)

const seed  = "http://www.zhenai.com/zhenghun"
func main() {

	engine.Run(engine.Request{
		Url: seed,
		ParserFunc:parser.ParseCityList,
	})
}
