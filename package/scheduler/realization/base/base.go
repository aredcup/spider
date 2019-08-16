package base

import (
	"github.com/golang/glog"
	"redcup/spider/package/scheduler/options"
	"redcup/spider/package/types"
)

type scheduler struct {
	requestChan chan *types.Request
	workChan    chan chan *types.Request

	requests []*types.Request
	works    []chan *types.Request
}

func NewBase(options *options.SchedulerOptions) *scheduler {
	return &scheduler{}
}

func (s *scheduler) AddRequest(request *types.Request) {
	s.requestChan <- request
}

func (s *scheduler) WorkerReady(worker chan *types.Request) {
	s.workChan <- worker
}

func (s *scheduler) GetWorkChan() chan *types.Request {
	return make(chan *types.Request)
}

func (s *scheduler) Run() {

	glog.Info("scheduler start")
	s.requestChan = make(chan *types.Request)
	s.workChan = make(chan chan *types.Request, 0)
	s.requests = make([]*types.Request, 0)
	s.works = make([]chan *types.Request, 0)

	go func() {

		for {
			var requestQ *types.Request
			var workQ chan *types.Request

			if len(s.requests) > 0 && len(s.works) > 0 {
				requestQ = s.requests[0]
				workQ = s.works[0]
			}

			select {
			case request := <-s.requestChan:
				s.requests = append(s.requests, request)
			case work := <-s.workChan:
				s.works = append(s.works, work)
			case workQ <- requestQ:
				s.requests = s.requests[1:]
				s.works = s.works[1:]
			}
		}
	}()
}
