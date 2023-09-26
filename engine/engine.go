package engine

import (
	"go-crawler/fetcher"
	"log"
)

func Run(seeds ...Request) {
	var requests []Request
	for _, e := range seeds {
		requests = append(requests, e)
	}
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		log.Printf("fetch url%s", r.Url)
		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("fetch errof :", err)
		}
		parseresult := r.ParseFunc(body)
		requests = append(requests, parseresult.Requests...)
		for _, items := range parseresult.Items {
			log.Printf("got items : %s\n", items)
		}
	}
}
