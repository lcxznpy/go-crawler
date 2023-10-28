package engine

type ParseResult struct {
	Requests []Request
	Items    []Item
}

type Item struct {
	Url     string
	Type    string
	Id      string
	Payload interface{}
}

type Parser interface {
	Parse(content []byte, url string) ParseResult
	Serialize() (name string, args interface{})
}

// 请求体，url和对应url的处理函数
type Request struct {
	Url   string
	Parse Parser
}

//func NilParse([]byte) ParseResult {
//	return ParseResult{}
//}

type NilParse struct {
}

func (NilParse) Parse(content []byte, url string) ParseResult {
	return ParseResult{}
}

func (NilParse) Serialize() (name string, args interface{}) {
	return "Nilparse", nil
}

type ParseFunc func(content []byte, url string) ParseResult

type FuncParser struct {
	parser ParseFunc
	name   string
}

func (f FuncParser) Parse(content []byte, url string) ParseResult {
	return f.parser(content, url)
}

func (f FuncParser) Serialize() (name string, args interface{}) {
	return f.name, nil
}

// 创建一个新的解析函数
func NewFuncParse(p ParseFunc, name string) *FuncParser {
	return &FuncParser{
		parser: p,
		name:   name,
	}
}
