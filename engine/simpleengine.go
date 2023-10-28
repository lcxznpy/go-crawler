package engine

import (
	"go-crawler/fetcher"
	"log"
)

type SimpleEngine struct {
	Scheduler Scheduler
}

func (s SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, e := range seeds {
		requests = append(requests, e)
	}
	// 开始是获取标签，后遍历每个标签的图书存入requesets中，最后获取图书详细内容
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		log.Printf("fetch url%s", r.Url)
		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("fetch errof :", err)
		}
		//获取douban图书的html body文件
		parseresult := r.ParseFunc(body)
		requests = append(requests, parseresult.Requests...)
		for _, items := range parseresult.Items {
			log.Printf("got items : %s\n", items)
		}
	}
}
