package client

import (
	"fmt"
	"polaris/crawler-distributed/config"
	"polaris/crawler-distributed/rpc"
	"polaris/crawler-distributed/worker"
	"polaris/crawler/engine"
)

func CreateProcessor() (engine.Processor, error) {
	client, err := rpcsupport.NewClient(fmt.Sprintf("%s", config.Worker0Host))
	if err != nil {
		return nil, err
	}

	return func(request engine.Request) (engine.ParseResult, error) {
		sReq := worker.SerializedRequest(request)
		var sResult worker.ParseResult
		err := client.Call(config.CrawlServiceRpc, sReq, &sResult)
		if err != nil {
			return engine.ParseResult{}, err
		}
		return worker.DeserializeParseResult(sResult), nil
	}, nil
}
