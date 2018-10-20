package main

import (
	"polaris/crawler-distributed/config"
	"polaris/crawler-distributed/rpc"
	"polaris/crawler-distributed/worker"
)

func main() {
	err := rpcsupport.ServeRpc(config.Worker0Host, worker.CrawlService{})
	if err != nil{
		panic(err)
	}
}
