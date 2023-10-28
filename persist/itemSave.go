package persist

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"go-crawler/engine"
	"log"
	"strings"
)

var ESClient *elastic.Client
var ESServerURL = []string{"http://127.0.0.1:9200"}

func ItemSave() (chan engine.Item, error) {
	out := make(chan engine.Item)

	ESClient, err := elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL(ESServerURL...))

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
			itemcount++
			Save(ESClient, item)

		}
	}()
	return out, nil
}

func Save(ESClient *elastic.Client, item engine.Item) error {
	//data, err := json.Marshal(item)
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//fmt.Println(bytes.NewReader(data))

	info, code, err := ESClient.Ping(strings.Join(ESServerURL, ",")).Do(context.Background())
	if err != nil {
		log.Fatalln("ping es failed", err.Error())
		return err
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	//bytes.NewReader(data)
	_, err = ESClient.Index().Index("crawler").BodyJson(item).Do(context.Background())
	if err != nil {
		return err
	}
	return nil

}
