package main

import (
	"github.com/golang/glog"
	clientoptions "redcup/spider/package/client/options"
	clientparser "redcup/spider/package/client/realization/parser"
	parserfetch "redcup/spider/package/parser"
	parseroptions "redcup/spider/package/parser/options"
	"redcup/spider/package/types"
	"redcup/spider/package/utils/config"
)

var nilParser types.Parser

func main() {

	baseOptions := new(types.BaseOptions)
	if err := config.Reload("./sources/spider.yaml", &baseOptions); err != nil {
		glog.Fatal(err)
	}

	nilParser = parserfetch.NewNilParser(parseroptions.NewParserOptions(baseOptions))

	client := clientparser.NewParser(clientoptions.NewClientOptions(baseOptions))
	client.RegisterFunc(types.RequestType_PARSER_URL, Parser)
	client.Run()
}

func Parser(request *types.Request, parser parserfetch.FetchURL) (result *types.ParseResult, err error) {

	content, err := parser.FetchURL(request.Host + request.Path)
	if err != nil {
		return
	}

	switch request.Parser {
	case types.ParserType_NIL:
		result, err = nilParser.Parse(content, request)
	}
	return
}
