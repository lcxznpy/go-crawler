package engine

import (
	"fmt"
	"go-crawler/fetcher"
	"go-crawler/scheduler"
	"log"
)

type BingFaEngine struct {
	Scheduler scheduler.SimpleScheduler
	WorkCount int
}

type Scheduler interface {
	Submit(Request)
	ConfigureWorkChan(chan Request)
	WorkReady(chan Request)
	Run()
}

func (b *BingFaEngine) Run(seeds ...Request) {
	in := make(chan Request)

	out := make(chan ParseResult)
	b.Scheduler.ConfigureWorkChan(in)
	for i := 0; i < b.WorkCount; i++ {
		CreateWorker(in, out)
	}
	for _, r := range seeds {
		b.Scheduler.Submit(r)
	}
	itemcount := 0
	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("got item:%d  %s", itemcount, item)
			itemcount++
		}
		for _, request := range result.Requests {
			b.Scheduler.Submit(request)
		}
	}
}

func CreateWorker(in chan Request, out chan ParseResult) {
	go func() {
		for {
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

func worker(r Request) (ParseResult, error) {
	fmt.Printf("fetch url : %s\n", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("fetch error:%s", err)
		return ParseResult{}, err
	}
	return r.ParseFunc(body), nil

}
