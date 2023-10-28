package main

import (
	"go-crawler/distributed/client"
	"go-crawler/distributed/config"
	client2 "go-crawler/distributed/worker/client"
	"go-crawler/engine"
	"go-crawler/parse"
	"go-crawler/scheduler"
)

func main() {
	itemsave, err := client.ItemSave(":1234")
	if err != nil {
		panic(err)
	}
	process, err := client2.CreateProcess()
	if err != nil {
		panic(err)
	}
	e := engine.BingFaEngine{
		Scheduler:        &scheduler.QueueScheduler{},
		WorkCount:        10,
		ItemChan:         itemsave,
		RequestProcessor: process,
	}
	e.Run(engine.Request{
		Url:   "https://book.douban.com/",
		Parse: engine.NewFuncParse(parse.ParseTag, config.ParseTagList),
	})
}
