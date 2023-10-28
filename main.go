package main

import (
	"go-crawler/engine"
	"go-crawler/parse"
	"go-crawler/persist"
	"go-crawler/scheduler"
)

func main() {
	//简易版run
	//e := engine.SimpleEngine{}
	//e.Run(engine.Request{
	//	Url:       "https://book.douban.com/",
	//	ParseFunc: parse.ParseTag,
	//})

	//并发
	//SimpleScheduler   单任务
	//QueueScheduler    队列调度器
	itemsave, err := persist.ItemSave()
	if err != nil {
		panic(err)
	}

	e := engine.BingFaEngine{
		Scheduler: &scheduler.QueueScheduler{},
		WorkCount: 10,
		ItemChan:  itemsave,
	}
	e.Run(engine.Request{
		Url:       "https://book.douban.com/",
		ParseFunc: parse.ParseTag,
	})
}
