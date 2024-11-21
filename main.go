package main

import (
	"log"
	"os"

	throughputLimiter "github.com/coderchirag/go-learning/concurrency/throughput_limiter"
	"github.com/coderchirag/go-learning/http"
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
			rateLimiterApp()
		case "server":
			http.StartServer()
	}
	
}
