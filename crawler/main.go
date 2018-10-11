package main

import (
	"polaris/crawler/engine"
	"polaris/crawler/persist"
	"polaris/crawler/scheduler"
	"polaris/crawler/zhenai/parser"
)

const seed = "http://www.zhenai.com/zhenghun"
const shanghai = "http://www.zhenai.com/zhenghun/shanghai"

func main() {

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    persist.ItemSaver(),
	}

	//e.Run(engine.Request{
	//	Url:        seed,
	//	ParserFunc: parser.ParseCityList,
	//})
	e.Run(engine.Request{
		Url:        shanghai,
		ParserFunc: parser.ParseCity,
	})
}
