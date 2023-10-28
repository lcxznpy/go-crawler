package client

import (
	"go-crawler/distributed/config"
	"go-crawler/distributed/rpcsupport"
	"go-crawler/distributed/worker"
	"go-crawler/engine"
)

func CreateProcess() (engine.Processor, error) {
	client, err := rpcsupport.NewClient(":1235")
	if err != nil {
		return nil, err
	}
	return func(req engine.Request) (engine.ParseResult, error) {
		sReq := worker.SerializeRequest(req)
		var sResult worker.ParseResult
		err := client.Call(config.CrawlServiceRpc, sReq, &sResult)
		if err != nil {
			return engine.ParseResult{}, err
		}
		return worker.DeserializeResult(sResult), nil
	}, nil
}
