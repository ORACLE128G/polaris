package client

import (
	"net/rpc"
	"polaris/crawler-distributed/config"
	"polaris/crawler-distributed/worker"
	"polaris/crawler/engine"
)

func CreateProcessor(
	pool chan *rpc.Client) (engine.Processor, error) {

	return func(request engine.Request) (engine.ParseResult, error) {
		sReq := worker.SerializedRequest(request)
		var sResult worker.ParseResult

		c := <-pool
		err := c.Call(config.CrawlServiceRpc, sReq, &sResult)
		if err != nil {
			return engine.ParseResult{}, err
		}
		return worker.DeserializeParseResult(sResult), nil
	}, nil
}
