package main

import (
	"log"
	"polaris/crawler-distributed/config"
	"polaris/crawler-distributed/rpc"
	"polaris/crawler-distributed/worker"
	"testing"
	"time"
)

func TestCrawlService(t *testing.T) {
	const host = ":8080"
	go rpcsupport.ServeRpc(
		host, worker.CrawlService{})
	time.Sleep(time.Second)

	client, err := rpcsupport.NewClient(
		host)
	if err != nil {
		panic(err)
	}

	request := worker.Request{
		Url: "http://album.zhenai.com/u/110591409",
		Parser: worker.SerializedParser{
			Name: config.ProfileParser,
			Args: "征婚",
		},
	}

	var result worker.ParseResult
	err = client.Call(config.CrawlServiceRpc, request, &result)
	if err != nil {
		t.Error(err)
	} else {
		log.Println(result)
	}

}
