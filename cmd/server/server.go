package main

import (
	"flag"
	"github.com/golang/glog"
	"math/rand"
	"redcup/spider/package/engine"
	engineoptions "redcup/spider/package/engine/options"
	"redcup/spider/package/types"
	"redcup/spider/package/utils/config"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	flag.Set("alsologtostderr", "true")
	flag.Parse()
	defer glog.Flush()

	baseOptions := new(types.BaseOptions)
	if err := config.Reload("./sources/spider.yaml", &baseOptions); err != nil {
		glog.Fatal(err)
	}
	glog.Infof("requests %+v", baseOptions.Requests)

	engineOptions := engineoptions.NewEngineOptions(baseOptions)
	engine := engine.NewEngine(engineOptions)
	engine.Run(getRequestsToRequestYaml(baseOptions.Requests)...)
}

func getRequestsToRequestYaml(requestYamls []*types.RequestYaml) (requests []*types.Request) {
	requests = make([]*types.Request, len(requestYamls))

	for i, requestYaml := range requestYamls {
		matchRules := make([]*types.MatchRule, len(requestYaml.MatchRules))
		for index, matchRule := range requestYaml.MatchRules {
			matchRules[index] = &types.MatchRule{
				Rule: matchRule.Rule,
				Keys: matchRule.Keys,
			}
		}

		requests[i] = &types.Request{
			Host:          requestYaml.Host,
			Path:          requestYaml.Path,
			Parser:        types.ParserType_Type(types.ParserType_Type_value[requestYaml.ParserName]),
			Type:          types.RequestType_PARSER_URL,
			PathRule:      requestYaml.PathRule,
			MatchRules:    matchRules,
			NeedMatchRule: requestYaml.NeedMatchPath,
		}
	}
	return
}
