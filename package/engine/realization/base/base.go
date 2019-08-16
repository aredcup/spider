package base

import (
	"github.com/golang/glog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"redcup/spider/package/engine/options"
	"redcup/spider/package/scheduler"
	"redcup/spider/package/storage"
	"redcup/spider/package/types"
	"strings"
	"sync"
)

type engine struct {
	scheduler        scheduler.Scheduler
	storage          storage.Storage
	clientResultChan chan *types.ParseResult
	visitedUrls      map[string]bool
	addr             string
	maxRequestCount  int

	muxtex sync.Mutex
}

func NewBase(options *options.EngineOptions) *engine {
	return &engine{
		scheduler:       options.Scheduler,
		maxRequestCount: options.MaxRequestCount,
		addr:            options.ListenAddr,
		storage:         options.Storage,
	}
}

// create worker
func (e *engine) Processor(stream types.Processor_ProcessorServer) error {
	glog.Info("server start")
	workChan := e.scheduler.GetWorkChan()
	e.scheduler.WorkerReady(workChan)
	go func(out chan *types.Request) {
		for {
			select {
			case request := <-out:
				glog.Info("send client request:", request)
				if err := stream.Send(request); err != nil {
					glog.Info("server close send")
					return
				}
			}
		}
	}(workChan)

	for {
		in, err := stream.Recv()
		if err != nil {
			goto ERR
		}
		glog.Info("get client result:", in)
		e.clientResultChan <- in
		e.scheduler.WorkerReady(workChan)
	}
ERR:
	glog.Info("server close")
	return nil
}

func (e *engine) AddRequest(request *types.Request) {

	e.muxtex.Lock()
	defer e.muxtex.Unlock()

	if !e.visitedUrls[request.Path] &&
		strings.Index(request.Path, request.PathRule) != -1 {
		e.scheduler.AddRequest(request)
		e.visitedUrls[request.Path] = true
	}
}

func (e *engine) Run(requests ...*types.Request) {
	glog.Info("engine start")
	e.clientResultChan = make(chan *types.ParseResult, 0)
	e.visitedUrls = make(map[string]bool, 0)

	e.scheduler.Run()
	e.storage.Run()

	go e.Listen()

	for _, request := range requests {
		e.AddRequest(request)
	}

	for {
		select {
		case clientResult := <-e.clientResultChan:
			e.FetchClientResult(clientResult)
		}
	}
}

func (e *engine) Listen() {

	listen, err := net.Listen("tcp", e.addr)
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()

	//注册grpc服务
	types.RegisterProcessorServer(grpcServer, e)

	// 在gRPC服务器上注册反射服务
	reflection.Register(grpcServer)

	err = grpcServer.Serve(listen)
	if err != nil {
		panic(err)
	}
}

func (e *engine) FetchClientResult(parseResult *types.ParseResult) {
	switch parseResult.Code {
	case types.ClientResultCode_ERROR:
		glog.Errorf("client error: %s", parseResult.ErrorMessage)
	case types.ClientResultCode_OK:

		if parseResult.Requests != nil && len(parseResult.Requests) > 0 {
			for _, request := range parseResult.Requests {
				e.AddRequest(request)
			}
		}

		if parseResult.Response != nil &&
			parseResult.Response.MatchRules != nil &&
			len(parseResult.Response.MatchRules) > 0 {
			e.storage.AddResponse(*parseResult.Response)
		}
	}
}
