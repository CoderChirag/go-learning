package worker_pools

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int){
	for j := range jobs {
		fmt.Println("Worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("Worker", id, "finished job", j)
		results <- j * 2
	}
}

func WorkerPoolsExample() {
	numJobs := 5

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w:=1; w<=3; w++ {
		go worker(w, jobs, results)
	}

	for j:=1; j<=numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for r:=1; r<=numJobs; r++ {
		<-results
	}
}