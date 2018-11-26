package engine

import "polaris/crawler/model"

type ParserFunc func(contents [] byte, url string) ParseResult

type Request struct {
	Url    string
	Parser Parser
}

type ParseResult struct {
	Requests [] Request
	Items    []Item
}

type Item struct {
	Url     string
	Id      string
	Type    string
	PayLoad model.Profile
	Anjuke  model.AnjukeProfile
}

type NilParser struct {
}

func (NilParser) Parse(_ []byte, _ string) ParseResult {
	return ParseResult{}
}

func (NilParser) Serialize() (name string, args interface{}) {
	return "NilParser", nil
}

type Parser interface {
	Parse(contents []byte, url string) ParseResult
	Serialize() (name string, args interface{})
}

type FuncParser struct {
	parser ParserFunc
	name   string
}

func (f *FuncParser) Parse(contents []byte, url string) ParseResult {
	return f.parser(contents, url)
}

func (f *FuncParser) Serialize() (name string, args interface{}) {
	return f.name, nil
}

func NewFuncParser(p ParserFunc, name string) *FuncParser {
	return &FuncParser{
		name:   name,
		parser: p,
	}
}
