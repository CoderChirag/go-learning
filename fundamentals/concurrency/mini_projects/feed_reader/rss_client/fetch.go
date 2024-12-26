package rss_client

import "github.com/coderchirag/go-learning/fundamentals/concurrency/mini_projects/feed_reader/types"

func Fetch(domain string) types.Fetcher {
	return &rssClient{url: domain}
}