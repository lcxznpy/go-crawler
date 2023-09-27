package engine

import (
	"go-crawler/engine"
	"go-crawler/parse"
)

func main() {
	//简易版run
	//engine.Run(engine.Request{
	//	Url:       "https://book.douban.com/",
	//	ParseFunc: parse.ParseTag,
	//})

	e := engine.BingFaEngine{
		Scheduler: &engine.SimpleScheduler{},
		WorkCount: 100,
	}
	e.Run(engine.Request{
		Url:       "https://book.douban.com/",
		ParseFunc: parse.ParseTag,
	})
}
