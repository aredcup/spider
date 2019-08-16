package parser

import (
	"fmt"
	"io/ioutil"
	"os"
	"redcup/spider/package/parser/options"
	"redcup/spider/package/types"
	"redcup/spider/package/utils/config"
	"testing"
)

func TestNilParser_Parse(t *testing.T) {
	baseOptions := new(types.BaseOptions)
	err := config.Reload("../../sources/spider.yaml", baseOptions)
	if err != nil {
		panic(err)
	}

	request :=  &types.Request{
		Host:       baseOptions.Requests[0].Host,
		Path:       baseOptions.Requests[0].Path,
		Parser:     types.ParserType_Type(types.ParserType_Type_value[baseOptions.Requests[0].ParserName]),
		Type:       types.RequestType_PARSER_URL,
		PathRule:   baseOptions.Requests[0].PathRule,
	}


	f, err := os.Open("../../regexp")
	if err != nil {
		panic(err)
	}

	b, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	parser := NewNilParser(options.NewParserOptions(&types.BaseOptions{}))



	clientResult, err := parser.Parse(b,request)
	if err != nil {
		panic(err)
	}
	fmt.Println(clientResult)
}

func TestFetch(t *testing.T) {
	content,err := Fetch("https://www.lagou.com/jobs/6282891.html?show=9eabea24f1cb4b44b294afad6625b1e8")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s",content)
}
