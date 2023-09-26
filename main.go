package main

import (
	"go-crawler/engine"
	"go-crawler/parse"
)

func main() {
	engine.Run(engine.Request{
		Url:       "https://book.douban.com/",
		ParseFunc: parse.ParseTag,
	})
}
