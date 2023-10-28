package engine

import (
	"go-crawler/fetcher"
	"log"
)

type Processor func(Request) (ParseResult, error)

type BingFaEngine struct {
	Scheduler        Scheduler //调度器类型
	WorkCount        int       //工人数量
	ItemChan         chan Item //存处理的结果
	RequestProcessor Processor //
}

type Scheduler interface {
	Submit(Request)
	WorkReady(chan Request)
	Run()
	WorkChan() chan Request //返回一个工作chan or多个工人chan
}

func (b *BingFaEngine) Run(seeds ...Request) {

	// 存输出结果
	out := make(chan ParseResult)

	//调度器工作
	b.Scheduler.Run()

	//创建工人，及分配并发调度器
	for i := 0; i < b.WorkCount; i++ {
		//创建工人，其实就是开启goroutine，监听请求的到来，处理请求
		b.CreateWorker(b.Scheduler.WorkChan(), out, b.Scheduler)
	}
	//提交任务进入任务队列
	for _, r := range seeds {
		b.Scheduler.Submit(r)
	}
	itemcount := 0
	//处理结果
	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("%d  %s", itemcount, item.Url)
			itemcount++
			go func() {
				b.ItemChan <- item
			}()
		}
		//将结果中的请求放入请求队列中
		for _, request := range result.Requests {
			b.Scheduler.Submit(request)
		}
	}
}

// 创建工人，
func (b *BingFaEngine) CreateWorker(in chan Request, out chan ParseResult, s Scheduler) {
	//开
	go func() {
		for {
			//每个工人都有自己的in chan用来接收并处理任务
			s.WorkReady(in)                            //告诉工人的接收队列已经空闲
			request := <-in                            //从工人的in  任务处理队列 中读 任务
			result, err := b.RequestProcessor(request) // 处理任务
			if err != nil {
				continue
			}
			//结果存入out中
			out <- result
		}
	}()
}

func Worker(r Request) (ParseResult, error) {
	//fmt.Printf("fetch url : %s\n", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("fetch error:%s", err)
		return ParseResult{}, err
	}
	return r.Parse.Parse(body, r.Url), nil

}
