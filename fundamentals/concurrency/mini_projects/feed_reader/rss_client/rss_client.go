package rss_client

import (
	"fmt"
	"net/http"
	"time"

	"github.com/coderchirag/go-learning/fundamentals/concurrency/mini_projects/feed_reader/types"
)

var _ types.Fetcher = (*rssClient)(nil)

type rssClient struct {
	url string
}

func (c *rssClient) Fetch() ([]string, time.Time, error) {
	res, err := http.Get(c.url)
	if err != nil {
		return nil, time.Now().Add(time.Second), err
	}
	b := make([]byte, 15)
	_, e := res.Body.Read(b)
	if e != nil {
		return nil, time.Now().Add(time.Second), err
	}

	return []string{fmt.Sprintf("[%s]", c.url), string(b)}, time.Now().Add(time.Second), nil
}