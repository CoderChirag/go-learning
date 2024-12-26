package feed_reader

import (
	"fmt"
	"time"

	"github.com/coderchirag/go-learning/fundamentals/concurrency/mini_projects/feed_reader/rss_client"
	"github.com/coderchirag/go-learning/fundamentals/concurrency/mini_projects/feed_reader/subscription"
)

func App(){
	// merged := subscription.Merge(
		merged := subscription.Subscribe(rss_client.Fetch("https://blob.golang.org"))
		// subscription.Subscribe(rss_client.Fetch("https://googleblog.blogspot.com")),
		// subscription.Subscribe(rss_client.Fetch("https://googledevelopers.blogspot.com")),
	// )

	time.AfterFunc(10*time.Second, func() {fmt.Println("Closed:", merged.Close())})

	for it := range merged.Updates() {
		fmt.Println(it)
	}

	panic("show me the stack")
}