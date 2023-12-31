package parse

import (
	"go-crawler/engine"
	"go-crawler/model"
	"regexp"
)

//var autoRe = regexp.MustCompile(`<span class="pl"> 作者</span>:[\d\D]*?<a.*?>([^<]+)</a>`)
//var scoreRe = regexp.MustCompile(`<strong class="ll rating_num " property="v:average">([^<]+)</strong>`)
//var publRe = regexp.MustCompile(`<span class="pl">出版社:</span>[\d\D]*?([^<]+)</a>`)
//var pageRe = regexp.MustCompile(`<span class="pl">页数:</span>([^<]+)<br/>`)
//var priceRe = regexp.MustCompile(`<span class="pl">定价:</span>([^<]+)<br/>`)
//var introRe = regexp.MustCompile(`<div class="intro">[\d\D]*?<p>([^<]+)</p></div>`)

var (
	authorRe    = regexp.MustCompile(`<span class="pl"> 作者</span>:[\d\D]*?<a.*?>([^<]+)</a>`)
	publisherRe = regexp.MustCompile(`<span class="pl">出版社:</span>[\d\D]*?<a.*?>([^<]+)</a>`)
	pagesRe     = regexp.MustCompile(`<span class="pl">页数:</span> ([^<]+)<br/>`)
	priceRe     = regexp.MustCompile(`<span class="pl">定价:</span> ([^<]+)<br/>`)
	scoreRe     = regexp.MustCompile(`<strong class="ll rating_num " property="v:average"> ([^<]+) </strong>`)
	introRe     = regexp.MustCompile(`<div class="intro">[\d\D]*?<p>([^<]+)</p></div>`)
)

func ParseBookDetail(content []byte, name string) engine.ParseResult {
	//fmt.Printf("%s", content)
	bookdetail := model.BookDetail{}
	bookdetail.BookName = name
	bookdetail.Author = ExtraString(content, authorRe)

	bookdetail.PageNum = ExtraString(content, pagesRe)

	bookdetail.Publisher = ExtraString(content, publisherRe)
	bookdetail.Info = ExtraString(content, introRe)
	bookdetail.Score = ExtraString(content, scoreRe)
	bookdetail.Price = ExtraString(content, priceRe)
	result := engine.ParseResult{
		Items: []interface{}{bookdetail},
	}
	return result
}

func ExtraString(content []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(content)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
