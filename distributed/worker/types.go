package worker

import (
	"errors"
	"fmt"
	"go-crawler/distributed/config"
	"go-crawler/engine"
	"go-crawler/parse"
	"log"
)

type SerializedParser struct {
	Name string
	Args interface{}
}

type Request struct {
	Url    string
	Parser SerializedParser
}

type ParseResult struct {
	Items    []engine.Item
	Requests []Request
}

// 将远端传来的request转成本地可以用的request
func SerializeRequest(r engine.Request) Request {
	name, args := r.Parse.Serialize()
	return Request{
		Url: r.Url,
		Parser: SerializedParser{
			Name: name,
			Args: args,
		},
	}
}

func SerializeResult(r engine.ParseResult) ParseResult {
	result := ParseResult{
		Items: r.Items,
	}
	for _, req := range r.Requests {
		result.Requests = append(result.Requests, SerializeRequest(req))
	}
	return result
}

func DeserializeRequest(r Request) (engine.Request, error) {
	parser, err := deserializeParser(r.Parser)
	if err != nil {
		return engine.Request{}, err
	}
	return engine.Request{
		Url:   r.Url,
		Parse: parser,
	}, nil
}

func DeserializeResult(r ParseResult) engine.ParseResult {
	result := engine.ParseResult{
		Items: r.Items,
	}
	for _, req := range r.Requests {
		engineReq, err := DeserializeRequest(req)
		if err != nil {
			log.Printf("error deserilizing request:%v", err)
			continue
		}
		result.Requests = append(result.Requests, engineReq)
	}

	return result
}

func deserializeParser(p SerializedParser) (engine.Parser, error) {
	switch p.Name {
	case config.ParseTagList:
		return engine.NewFuncParse(
			parse.ParseTag,
			config.ParseTagList), nil
	case config.ParseBookList:
		return engine.NewFuncParse(
			parse.ParseBookList,
			config.ParseBookList), nil
	case config.NilParser:
		return &engine.NilParse{}, nil
	case config.ParseBookDetail:
		if name, ok := p.Args.(string); ok {
			return parse.NewBookDetailParser(name), nil
		} else {
			return nil, fmt.Errorf("invalid arg:%v", p.Args)
		}
	default:
		return nil, errors.New("unknown parser name")
	}
}
