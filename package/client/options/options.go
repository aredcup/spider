package options

import (
	"redcup/spider/package/parser"
	"redcup/spider/package/parser/realization/lagou"
	"redcup/spider/package/types"
)

type ClientOptions struct {
	MasterAddr string
	LimitTime  uint64
	Fetch      parser.FetchURL
}

func NewClientOptions(baseOptions *types.BaseOptions) *ClientOptions {
	return &ClientOptions{
		MasterAddr: baseOptions.ClientOptions.MasterAddr,
		LimitTime:  baseOptions.ClientOptions.LimitTime,
		Fetch:      lagou.NewLagouFetch(),
	}
}
