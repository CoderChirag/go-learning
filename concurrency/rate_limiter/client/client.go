package client

import (
	"fmt"
	"time"

	"github.com/coderchirag/go-learning/concurrency/rate_limiter/transport"
)

type Client interface {
	New() Client
	SendRequest(args []any, fn func(args... any) any, resultChan chan any)
}

type RPCClient struct {}

func (client *RPCClient) New () Client {
	return client
}

func (client *RPCClient) SendRequest(args []any, fn func(args... any) any, resultChan chan any) {
	dto := transport.Request[any, any]{Args: args, Fn: fn, ResultChan: resultChan}
	fmt.Println("Sending Request with args:", args, time.Now())
	transport.RequestsQueue <- &dto
	fmt.Println("Received Response for args:", args, <-resultChan, time.Now())
}