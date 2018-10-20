package main

import (
	"polaris/crawler-distributed/rpc"
	"polaris/crawler-distributed/worker"
)

func main() {
	rpcsupport.ServeRpc(":8080", worker.CrawlService{})
}
