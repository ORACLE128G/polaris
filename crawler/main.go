package main

import (
	"polaris/crawler-distributed/config"
	"polaris/crawler/anjuke/parser"
	"polaris/crawler/engine"
	"polaris/crawler/persist"
	"polaris/crawler/scheduler"
)

const seed = "http://www.zhenai.com/zhenghun"
const shanghai = "http://www.zhenai.com/zhenghun/shanghai"
const anjukeCityList = "https://www.anjuke.com/sy-city.html"

func main() {

	// unblock the fllowing code, we can got anjuke.com all of pages.
	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         persist.ItemSaver(),
		RequestProcessor: engine.AnjukeWorker,
	}

	e.Run(engine.Request{
		Url: anjukeCityList,
		Parser: engine.NewFuncParser(
			parser.ParseCityList,
			config.ParseCityList),
	})

	// for zhenai.com crawler.
	/*e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         persist.ItemSaver(),
		RequestProcessor: engine.Worker,
	}

	e.Run(engine.Request{
		Url: seed,
		Parser: engine.NewFuncParser(
			parser.ParseCityList,
			config.ParseCityList),
	})*/

	// unblock the following code, we can got zhenai.com user details.

	/*e.Run(engine.Request{
		Url: shanghai,
		Parser: engine.NewFuncParser(
			parser.ParseCity,
			"ParseCity"),
	})*/
}
