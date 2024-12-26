package subscription

import (
	"github.com/coderchirag/go-learning/fundamentals/concurrency/mini_projects/feed_reader/types"
)

func Subscribe(fetcher types.Fetcher) types.Subscription {
	s := &subscription{fetcher: fetcher, updates: make(chan string), closing: make(chan chan error)}
	go s.loop()
	return s
}

// func Merge(subs ...types.Subscription) types.Subscription {
// 	// for sub := range subs {

// 	// }
// }