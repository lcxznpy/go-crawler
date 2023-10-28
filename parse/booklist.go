package parse

import (
	"go-crawler/engine"
	"regexp"
)

const BookListRe = `<a href="([^"]+)" title="([^"]+)"`

func ParseBookList(content []byte, _ string) engine.ParseResult {
	//fmt.Printf("%s", content)
	re := regexp.MustCompile(BookListRe)
	match := re.FindAllSubmatch(content, -1)
	result := engine.ParseResult{}
	cnt := 0
	for _, m := range match {
		bookname := string(m[2])
		//result.Items = append(result.Items, engine.Item{
		//	Url: string(m[2]),
		//})
		cnt++
		if cnt > 100 {
			return result
		}
		url := string(m[1])
		result.Requests = append(result.Requests, engine.Request{
			Url: url,
			//函数式编程解决，跳转界面无书名问题
			//ParseFunc: func(c []byte) engine.ParseResult {
			//
			//	return ParseBookDetail(c, bookname, string(m[1]))
			//},
			Parse: NewBookDetailParser(bookname),
		})
	}
	return result

}
