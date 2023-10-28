package server

import (
	"go-crawler/distributed/worker"
	"go-crawler/engine"
)

type CrawlService struct {
}

func (CrawlService) Process(req worker.Request, result *worker.ParseResult) error {
	engineReq, err := worker.DeserializeRequest(req)
	if err != nil {
		return err
	}
	//当前节点执行工作
	engineResult, err := engine.Worker(engineReq)
	if err != nil {
		return err
	}
	*result = worker.SerializeResult(engineResult)
	return nil
}
