package parse

import (
	"go-crawler/engine"
	"regexp"
)

const BookListRe = `<a href="([^"]+)" title="([^"]+)"`

func ParseBookList(content []byte) engine.ParseResult {
	//fmt.Printf("%s", content)
	re := regexp.MustCompile(BookListRe)
	match := re.FindAllSubmatch(content, -1)
	result := engine.ParseResult{}

	for _, m := range match {
		bookname := string(m[2])
		result.Items = append(result.Items, m[2])
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			//函数式编程解决，跳转界面无书名问题
			ParseFunc: func(c []byte) engine.ParseResult {
				return ParseBookDetail(c, bookname)
			},
		})
	}
	return result

}
