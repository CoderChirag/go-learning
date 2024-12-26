package main

import (
	"log"
	"os"

	rateLimiterApp "github.com/coderchirag/go-learning/fundamentals/concurrency/mini_projects/rate_limiter/app"
	throughputLimiter "github.com/coderchirag/go-learning/fundamentals/concurrency/mini_projects/throughput_limiter"
	"github.com/coderchirag/go-learning/http"
	"github.com/coderchirag/go-learning/playground"
)


func main() {
	defer func(){
		err := recover(); 
		if err != nil && len(os.Args) == 1 {
			log.Println("Please Specify Arguments for starting the application")
		}else if err != nil {
			panic(err)
		}
	}()

	app := os.Args[1]
	switch app {
		case "tpl":
			throughputLimiter.DoTaskConcurrentlyWithThroughputLimit(20, 3)
		case "rl":
			rateLimiterApp.RateLimiterApp()
		case "server":
			http.StartServer()
		case "pg":
			playground.Playground()
	}

}