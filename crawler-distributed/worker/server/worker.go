package main

import (
	"flag"
	"fmt"
	"log"
	"polaris/crawler-distributed/rpc"
	"polaris/crawler-distributed/worker"
)

var port = flag.Int("port", 0,
	"The port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		log.Printf("Must specify a port ")
		return
	}
	err := rpcsupport.ServeRpc(fmt.Sprintf(
		":%d", *port), worker.CrawlService{})
	if err != nil {
		panic(err)
	}
}
