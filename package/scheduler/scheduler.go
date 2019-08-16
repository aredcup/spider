package scheduler

import (
	"redcup/spider/package/scheduler/options"
	"redcup/spider/package/scheduler/realization/base"
	"redcup/spider/package/types"
)

type Scheduler interface {
	AddRequest(request *types.Request)
	WorkerReady(w chan *types.Request)
	GetWorkChan() chan *types.Request
	Run()
}

func NewScheduler(options *options.SchedulerOptions) Scheduler {
	return base.NewBase(options)
}
