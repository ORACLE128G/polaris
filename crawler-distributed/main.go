package main

import (
	"flag"
	"log"
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

	if *itemSaverHost == "" && *workerHosts == "" {
		log.Println("Must specify one or more itemsaver host, like this:\n" +
			"go run main.go --itemsaver_host=\":9200\"")
		log.Println("Must specify one or more worker host, like this:\n" +
			"go run main.go --worker_host=\"9300\"")
		log.Print("You can typing the go run main.go --help for get help")
		return
	}
	if *itemSaverHost == "" {
		log.Print("Must specify one or more itemsaver host, like this:\n" +
			"go run main.go --itemsaver_host=\":9200\"")
		return
	}

	if *workerHosts == "" {
		log.Print("Must specify one or more worker host, like this:\n" +
			"go run main.go --worker_host=\"9300\"")
		return
	}

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

