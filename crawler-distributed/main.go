package main

import (
	"polaris/crawler-distributed/config"
	itemsaver "polaris/crawler-distributed/persist/client"
	worker "polaris/crawler-distributed/worker/client"
	"polaris/crawler/engine"
	"polaris/crawler/scheduler"
	"polaris/crawler/zhenai/parser"
)

func main() {
	// run condition:
	// run itemsaver.go first
	// then to run worker.go
	itemChan, err := itemsaver.ItemSaver(config.ItemSaver0Host)
	if err != nil {
		panic(err)
	}

	processor, err := worker.CreateProcessor()
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: processor,
	}

	e.Run(engine.Request{
		Url:    "http://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(parser.ParseCityList, config.ParseCityList),
	})
}
