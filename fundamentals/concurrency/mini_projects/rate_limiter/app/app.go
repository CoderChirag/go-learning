package app

import (
	"log"
	"math/rand"
	"sync"

	clientLib "github.com/coderchirag/go-learning/fundamentals/concurrency/mini_projects/rate_limiter/client"
	serverLib "github.com/coderchirag/go-learning/fundamentals/concurrency/mini_projects/rate_limiter/server"
	transport "github.com/coderchirag/go-learning/fundamentals/concurrency/mini_projects/rate_limiter/transport"
)

func RateLimiterApp(){
	var server serverLib.RPCServer
	var client clientLib.RPCClient

	log.Println("Starting Rate Limiter Application...")

	runServer(&server)

	wg := runClient(&client)
	defer wg.Wait()
}

func runServer(server serverLib.Server) {
	go func(){
		server.Serve(transport.RequestsQueue)
	}()
}

func runClient(client clientLib.Client) *sync.WaitGroup {
	var wg sync.WaitGroup
	resultChan := make(chan any)

	for i:=0; i<10; i++ {
		wg.Add(1)
		go func() {
			args := []any{rand.Intn(10), rand.Intn(10), rand.Intn(10)}
			client.New().SendRequest(args, sum, resultChan)
			wg.Done()
		}()
	}
	return &wg
}