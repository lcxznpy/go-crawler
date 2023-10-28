package parse

import (
	"go-crawler/distributed/config"
	"go-crawler/engine"
	"regexp"
)

const regexpStr = `<a href="([^"]+)" class="tag">([^"]+)</a>`

// 通过正则获取网页标签
func ParseTag(content []byte, _ string) engine.ParseResult {
	//<a href="/tag/小说" class="tag">小说</a>
	re := regexp.MustCompile(regexpStr)
	match := re.FindAllSubmatch(content, -1)

	result := engine.ParseResult{}
	cnt := 0
	for _, m := range match {
		//result.Items = append(result.Items, engine.Item{
		//	Url: string(m[2]),
		//})
		cnt++
		if cnt > 10 {

			return result
		}
		result.Requests = append(result.Requests, engine.Request{
			Url:   "https://book.douban.com" + string(m[1]),
			Parse: engine.NewFuncParse(ParseBookList, config.ParseBookList),
		})
	}
	return result
	//for _, m := range match {
	//	fmt.Printf("m[0]:%s,m[1]:%s,m[2]:%s\n", m[0], m[1], m[2])
	//}
}
