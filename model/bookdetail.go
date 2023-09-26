package model

type BookDetail struct {
	BookName  string //书名
	Author    string //作者
	Publisher string //出版社
	PageNum   string //页数
	Price     string //价格
	Score     string //豆瓣评分
	Info      string //内容简介
}

func (b BookDetail) String() string {
	return "书名:" + b.BookName + "作者:" + b.Author + "出版社:" + b.Publisher + "页数:" + b.PageNum + "价格:" + b.Price + "豆瓣评分:" + b.Score + "简介:" + b.Info
}
