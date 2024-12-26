package rate_limiter

import (
	"sync"
	"time"
)

func WithRateLimit[T any](fn func(args ...T), duration time.Duration, limit int) (func(args ...T), *sync.WaitGroup) {
	var wg sync.WaitGroup
	var mutex sync.Mutex
	lastReset := time.Now()
	tokens := limit

	return func(args ...T) {
		mutex.Lock()
		now := time.Now()
		if now.Sub(lastReset) >= duration {
			tokens = limit
			lastReset = now
		}

		for tokens <= 0 {
			mutex.Unlock()
			time.Sleep(time.Millisecond * 100)

			mutex.Lock()
			now = time.Now()
			if now.Sub(lastReset) >= duration {
				tokens = limit
				lastReset = now
			}
		}

		tokens--
		mutex.Unlock()
		wg.Add(1)
		go func() {
			defer wg.Done()
			fn(args...)
		}()
	}, &wg
}