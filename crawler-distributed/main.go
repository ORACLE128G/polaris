package main

import (
	"flag"
	"polaris/crawler-distributed/config"
	itemsaver "polaris/crawler-distributed/persist/client"
	"polaris/crawler-distributed/pool"
	worker "polaris/crawler-distributed/worker/client"
	"polaris/crawler/engine"
	"polaris/crawler/scheduler"
	"polaris/crawler/zhenai/parser"
	"strings"
)

var (
	itemSaverHost = flag.String("itemsaver_host", "",
		"itemsaver_host")

	workerHosts = flag.String("worker_host", "",
		"worker_host(comma separated)")
)
func main() {
	// run condition:
	// run itemsaver.go first
	// then to run worker.go
	flag.Parse()
	itemChan, err := itemsaver.ItemSaver(*itemSaverHost)
	if err != nil {
		panic(err)
	}

	clientPool := pool.CreateClientPool(strings.Split(*workerHosts, ","))
	processor, err := worker.CreateProcessor(clientPool)
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

