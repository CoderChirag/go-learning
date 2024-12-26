package mutexes

import (
	"fmt"
	"sync"
)

type container struct {
	mu sync.Mutex
	counters map[string]int
}

func (c *container) inc(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

func MutexExample(){
	c := container{
		counters: map[string]int{"a":0, "b": 0},
	}

	var wg sync.WaitGroup

	increase := func(name string, n int){
		for i:=0; i<n; i++ {
			c.inc(name)
		}
		wg.Done()
	}

	wg.Add(3)
	go increase("a", 10000)
	go increase("a", 10000)
	go increase("b", 10000)

	wg.Wait()
	fmt.Println(c.counters)
}