package client

import (
	"go-crawler/distributed/config"
	"go-crawler/distributed/rpcsupport"
	"go-crawler/engine"
	"log"
)

func ItemSave(host string) (chan engine.Item, error) {
	out := make(chan engine.Item)

	//初始化rpc客户端
	client, err := rpcsupport.NewClient(host)

	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
		panic(err)
	}

	//函数式编程，在goroutine中使用了外部的变量，那么这个变量就可以一直用下去
	go func() {
		itemcount := 0
		for {
			item := <-out
			log.Printf("item saver got:%d  %s", itemcount, item)
			result := ""
			err := client.Call(config.ItemSaverRpc, item, &result)
			if err != nil {
				log.Printf("item saver got error :%v", err)
			}
			itemcount++

		}
	}()
	return out, nil
}
