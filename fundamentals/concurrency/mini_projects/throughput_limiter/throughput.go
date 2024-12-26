package throughput_limiter

import (
	"fmt"
	"sync"
	"time"
)

func task(id int, args... any) {
	fmt.Println("Handler", id, "started", "args:", args, time.Now())
	time.Sleep(3 * time.Second)
	fmt.Println("Handler", id, "finished", time.Now())
}

func DoTaskConcurrentlyWithThroughputLimit(count int, throughputLimit int) {
	taskFn, wg := WithThroughputLimit(task, throughputLimit)
	defer wg.Wait()
	for i:=0; i<count; i++ {
		taskFn(i, 1, 2, "abcd")
	}
}

func WithThroughputLimit[T any](fn func(id int, args... T), limit int) (func(id int, args... T), *sync.WaitGroup) {
	var wg sync.WaitGroup
	sem := make(chan int, limit)
	return func(id int, args... T) {
		sem <- id
		wg.Add(1)
		fmt.Println("Handler", id, "scheduled", "args:", args, time.Now())
		go func() {
			defer wg.Done()
			fn(id, args...)
			<- sem
		}()
	}, &wg
}
