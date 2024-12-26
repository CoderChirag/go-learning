package server

import (
	"time"

	"github.com/coderchirag/go-learning/fundamentals/concurrency/mini_projects/rate_limiter"
	"github.com/coderchirag/go-learning/fundamentals/concurrency/mini_projects/rate_limiter/transport"
)

type Server interface {
	Serve(reqQueue chan *transport.Request[any, any])
}

type RPCServer struct {}

func handle(req *transport.Request[any, any]) {
	// fmt.Println("Received request with args:", req.Args, time.Now())
	req.ResultChan <- req.Fn(req.Args...)
}

func (server *RPCServer) Serve(reqQueue chan *transport.Request[any, any]) {
	wrapper := func(args ...any) {
		handle(args[0].(*transport.Request[any, any]))
	}
	handler, _ := rate_limiter.WithRateLimit(wrapper, time.Second, 3)

	for req := range reqQueue {
		go func(){
			handler(req)
		}()
	}
}