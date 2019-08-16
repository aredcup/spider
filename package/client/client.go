package client

import (
	"redcup/spider/package/parser"
	"redcup/spider/package/types"
)

type HandlerFunc func(request *types.Request, parser parser.FetchURL) (result *types.ParseResult, err error)

type Client interface {
	Run()
	RegisterFunc(path types.RequestType_Type, fn HandlerFunc)
	Close()
}
