package main

import (
	"fmt"
	"github.com/olivere/elastic/v7"
	"go-crawler/distributed/persist"
	"go-crawler/distributed/rpcsupport"
)

var ESServerURL = []string{"http://127.0.0.1:9200"}

func main() {
	serveRpc(":1234")
}

// 初始化els客户端，建立rpc服务端
func serveRpc(host string) error {
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL(ESServerURL...))
	if err != nil {
		return err
	}
	fmt.Println("success connect els")
	return rpcsupport.ServerRpc(host, &persist.ItemService{
		Client: client,
	})
}
