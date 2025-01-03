package stateful_goroutines

import (
	"fmt"
	"math/rand/v2"
	"sync/atomic"
	"time"
)

type readOp struct {
	key int
	resp chan int
}

type writeOp struct {
	key int
	val int
	resp chan bool
}

func stateManager(reads <-chan readOp, writes <-chan writeOp){
	state := make(map[int]int)
	for {
		select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
		}
	}
}

func StatefulGoroutine(){
	var readOps uint64
	var writeOps uint64

	reads := make(chan readOp)
	writes := make(chan writeOp)

	go stateManager(reads, writes)

	for r:=0; r<100; r++ {
		go func(){
			for {
				read := readOp{key: rand.IntN(5), resp: make(chan int)}
				reads <- read
				<-read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for w:=0; w<10; w++ {
		go func(){
			for {
				write := writeOp{key: rand.IntN(5), val: rand.IntN(100), resp: make(chan bool)}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)

	fmt.Println("ReadOps:", atomic.LoadUint64(&readOps))
	fmt.Println("WriteOps:", atomic.LoadUint64(&writeOps))
}