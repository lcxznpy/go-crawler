package engine

type ParseResult struct {
	Requests []Request
	Items    []interface{}
}

// 请求体，url和对应url的处理函数
type Request struct {
	Url       string
	ParseFunc func([]byte) ParseResult
}

func NilParse([]byte) ParseResult {
	return ParseResult{}
}
