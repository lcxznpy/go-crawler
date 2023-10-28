package persist

import (
	"github.com/olivere/elastic/v7"
	"go-crawler/engine"
	"go-crawler/persist"
)

// 持久化服务
type ItemService struct {
	Client *elastic.Client
}

// 持久化
func (s *ItemService) Save(item engine.Item, result *string) error {
	err := persist.Save(s.Client, item)
	if err == nil {
		*result = "ok"
	}
	return err
}
