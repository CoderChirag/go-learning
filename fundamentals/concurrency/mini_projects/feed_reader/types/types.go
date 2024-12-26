package types

import "time"

type Fetcher interface {
	Fetch() ([]string, time.Time, error)
}

type Subscription interface {
	Updates() <-chan string
	Close() error
}