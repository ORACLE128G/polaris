package worker

import (
	"errors"
	"fmt"
	"github.com/gpmgo/gopm/modules/log"
	"polaris/crawler-distributed/config"
	"polaris/crawler/engine"
	"polaris/crawler/zhenai/parser"
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

func SerializedRequest(r engine.Request) Request {
	name, args := r.Parser.Serialize()
	return Request{
		Url: r.Url,
		Parser: SerializedParser{
			Name: name,
			Args: args,
		},
	}
}

func SerializedParseResult(r engine.ParseResult) ParseResult {
	result := ParseResult{
		Items: r.Items,
	}
	for _, e := range r.Requests {
		result.Requests = append(result.Requests, SerializedRequest(e))
	}
	return result
}

func DeserializeRequest(r Request) (engine.Request, error) {
	p, err := deserializeParser(r.Parser)
	if err != nil {
		return engine.Request{}, err
	}
	return engine.Request{
		Url:    r.Url,
		Parser: p,
	}, nil
}

func DeserializeParseResult(r ParseResult) engine.ParseResult {
	result := engine.ParseResult{
		Items: r.Items,
	}

	for _, e := range r.Requests {
		request, err := DeserializeRequest(e)
		if err != nil {
			log.Error("error deserializing request: %v", err)
			continue
		}
		result.Requests = append(result.Requests, request)
	}
	return result
}

func deserializeParser(p SerializedParser) (engine.Parser, error) {
	switch p.Name {
	case config.ParseCity:
		return engine.NewFuncParser(
			parser.ParseCity,
			config.ParseCity), nil
	case config.ParseCityList:
		return engine.NewFuncParser(
			parser.ParseCityList,
			config.ParseCityList,
		), nil
	case config.ProfileParser:
		if userName, ok := p.Args.(string); ok {
			return parser.NewProfileParser(userName), nil
		} else {
			return nil, fmt.Errorf("invalid arg: %v", userName)
		}

	default:
		return nil, errors.New("unknown parser name")
	}
}