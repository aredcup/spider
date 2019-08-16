package engine

import (
	"redcup/spider/package/engine/options"
	"redcup/spider/package/engine/realization/base"
	"redcup/spider/package/types"
)

type Engine interface {
	AddRequest(request *types.Request)
	Run(requests ...*types.Request)
	Listen()
	types.ProcessorServer
}

type ReadyNotifier interface {
	WorkerReady(w chan *types.Request)
}

func NewEngine(options *options.EngineOptions) Engine {
	return base.NewBase(options)
}

