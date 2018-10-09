package main

import (
	"polaris/crawler/engine"
	"polaris/crawler/scheduler"
	"polaris/crawler/zhenai/parser"
)

const seed = "http://www.zhenai.com/zhenghun"

func main() {

	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.SimpleScheduler{

		},
		WorkerCount: 100,
	}

	e.Run(engine.Request{
		Url:        seed,
		ParserFunc: parser.ParseCityList,
	})
}
