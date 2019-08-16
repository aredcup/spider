package parser

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"redcup/spider/package/client"
	"redcup/spider/package/client/options"
	parserfetch "redcup/spider/package/parser"
	"redcup/spider/package/types"
	"sync"
	"time"
)

type parser struct {
	masterAddr string
	funcMap    map[types.RequestType_Type]client.HandlerFunc
	stream     types.Processor_ProcessorClient
	inChan     chan *types.Request
	outChan    chan *types.ParseResult
	closeChan  chan struct{}
	isClose    bool
	mux        sync.Mutex
	limitTime  uint64
	fetch      parserfetch.FetchURL
}

func NewParser(options *options.ClientOptions) *parser {
	return &parser{
		masterAddr: options.MasterAddr,
		limitTime:  options.LimitTime,
		outChan:    make(chan *types.ParseResult, 0),
		fetch:      options.Fetch,
		inChan:     make(chan *types.Request, 0),
		funcMap:    make(map[types.RequestType_Type]client.HandlerFunc, 0),
	}
}

func (c *parser) RegisterFunc(path types.RequestType_Type, fn client.HandlerFunc) {
	c.funcMap[path] = fn
}

func (c *parser) Close() {
	fmt.Println("client start to close")
	if c.isClose {
		return
	}

	c.mux.Lock()
	defer c.mux.Unlock()

	close(c.closeChan)
	close(c.inChan)
	close(c.outChan)
	c.isClose = true

	fmt.Println("client close")
}

func (c *parser) Run() {
	var (
		err    error
		conn   *grpc.ClientConn
		client types.ProcessorClient
	)
	conn, err = grpc.Dial(c.masterAddr, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	// 函数结束时关闭连接
	defer conn.Close()

	// 创建Waiter服务的客户端
	client = types.NewProcessorClient(conn)
	c.stream, err = client.Processor(context.Background())
	if err != nil {
		panic(err)
	}

	go c.readLoop()
	go c.writeLoop()

	for {
		if c.isClose {
			break
		}
		select {
		case request := <-c.inChan:
			fmt.Println("request:", request)
			result, err := c.funcMap[request.Type](request, c.fetch)
			fmt.Println("result:", result)
			if result == nil {
				result = &types.ParseResult{}
			}

			result.Requested = request
			if err != nil {
				result.Code = types.ClientResultCode_ERROR
				result.ErrorMessage = err.Error()
				fmt.Println(err)
			}

			c.outChan <- result
		case <-c.closeChan:
			c.Close()
		}
	}

}

func (c *parser) readLoop() {
	fmt.Println("client readLoop:")
	for {
		time.Sleep(time.Duration(c.limitTime) * time.Millisecond)
		request, err := c.stream.Recv()
		if err != nil && err != io.EOF {
			c.closeChan <- struct{}{}
		}
		c.inChan <- request
	}
}

func (c *parser) writeLoop() {
	fmt.Println("client writeLoop:")
	for {
		select {
		case result := <-c.outChan:
			if err := c.stream.Send(result); err != nil {
				c.closeChan <- struct{}{}
			}
		case <-c.closeChan:
			c.Close()
		}
	}
}
