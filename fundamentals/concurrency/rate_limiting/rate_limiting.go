package rate_limiting

import (
	"fmt"
	"time"
)

/* This rate limiter would process a request every 200ms */
func SimpleRateLimiter(){
	requests := make(chan int, 5)
	for i:=1; i<=5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(200 * time.Millisecond)

	for req := range requests {
		<-limiter
		fmt.Println("Served Request", req, time.Now())
	}
}

/* This rate limiter can process a burst of up to 3 requests every 200ms */
func BurstyRateLimiter(){
	requests := make(chan int, 5)
	for i:=1; i<=5; i++ {
		requests <- i
	}
	close(requests)

	burstyLimiter := make(chan time.Time, 3)
	for i:=0; i<3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(200 * time.Millisecond){
			burstyLimiter <- t
		}
	}()

	for req := range requests {
		<-burstyLimiter
		fmt.Println("Served Request", req, time.Now())
	}
}
