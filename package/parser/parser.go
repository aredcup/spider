package parser

import (
	"redcup/spider/package/parser/options"
	"redcup/spider/package/types"
	"regexp"
)

func NewNilParser(options *options.ParserOptions) types.Parser {
	return &nilParser{}
}

type nilParser struct {
	isInit bool
}

func (p *nilParser) Parse(contents []byte, request *types.Request) (parseResult *types.ParseResult, err error) {
	if !p.isInit {
		p.InitRegexp(request)
	}

	parseResult = &types.ParseResult{
		Requests: make([]*types.Request, 0),
		Response: &types.Response{},
	}

	for rule, regexp := range types.RegexpMap {
		matches := regexp.FindAllSubmatch(contents, -1)
		for _, match := range matches {
			if rule == "a[href]" {
				parseResult.Requests = append(parseResult.Requests, &types.Request{
					Host:       request.GetHost(),
					Path:       string(match[1]),
					MatchRules: request.GetMatchRules(),
					Parser:     request.GetParser(),
					Type:       types.RequestType_PARSER_URL,
					PathRule:   request.PathRule,
				})
			} else {

				//for _, matchRule := range request.GetMatchRules() {
				//	if rule == matchRule.Rule {
				//		if strings.Contains(matchRule.Keys, ",") {
				//			keys := strings.Split(matchRule.Keys, ",")
				//			for i := range keys {
				//				matchRule.Values += string(match[i+1]) + ","
				//			}
				//		} else {
				//			matchRule.Values = string(match[1])
				//		}
				//	}
				//}
				//
				//if request.MatchRules != nil {
				//	parseResult.Response.MatchRules = request.GetMatchRules()
				//}
			}
		}
	}

	return
}

func (p *nilParser) InitRegexp(request *types.Request) (err error) {

	for _, matchRule := range request.GetMatchRules() {
		if types.RegexpMap[matchRule.Rule] == nil {
			types.RegexpMap[matchRule.Rule] = regexp.MustCompile(matchRule.Rule)
		}
	}

	p.isInit = true
	return
}

func (*nilParser) Serialize() (name string, args interface{}) {
	panic("implement me")
}
