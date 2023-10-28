package main

import (
	"go-crawler/distributed/rpcsupport"
	"go-crawler/distributed/worker/server"
	"log"
)

// 建立工人服务端
func main() {
	//注册服务
	log.Fatal(rpcsupport.ServerRpc(":1235", &server.CrawlService{}))
}
