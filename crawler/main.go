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
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         persist.ItemSaver(),
		RequestProcessor: engine.Worker,
	}

	/*e.Run(engine.Request{
		Url: seed,
		Parser: engine.NewFuncParser(
			parser.ParseCityList,
			config.ParseCityList),
	})*/
	e.Run(engine.Request{
		Url: shanghai,
		Parser: engine.NewFuncParser(
			parser.ParseCity,
			"ParseCity"),
	})
}
