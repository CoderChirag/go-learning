package subscription

import (
	"time"

	"github.com/coderchirag/go-learning/fundamentals/concurrency/mini_projects/feed_reader/types"
)

var _ types.Subscription = (*subscription)(nil)

type fetchResult struct {
	fetched []string
	next time.Time
	err error
}

type subscription struct {
	fetcher types.Fetcher
	updates chan string
	closing chan chan error
}

func (s *subscription) Updates() <-chan string {
	return s.updates
}

func (s *subscription) Close() error {
	errc := make(chan error)
	s.closing <-errc
	return <-errc
}

func (s *subscription) loop() {
	var pending []string
	var next time.Time
	var err error
	var fetchDone chan fetchResult
	for {
		var fetchDelay time.Duration
		var first string
		var updates chan string
		var startFetch <-chan time.Time
		if now := time.Now(); next.After(now) {
			fetchDelay = next.Sub(now)
		}
		if fetchDone == nil {
			startFetch = time.After(fetchDelay)
		}

		if len(pending) > 0 {
			first = pending[0]
			updates = s.updates
		}

		select {
			case errc := <-s.closing:
				errc <-err
				close(s.updates)
				return
			case <-startFetch:
					fetchDone = make(chan fetchResult, 1)
					go func(){
						fetched, next, err := s.fetcher.Fetch()
						fetchDone <- fetchResult{fetched, next, err}
					}()
			case result := <-fetchDone:
				fetchDone = nil
				pending = append(pending, result.fetched...)
				next = result.next
				err = result.err
			case updates <-first:
				pending = pending[1:]
		}
	}
}