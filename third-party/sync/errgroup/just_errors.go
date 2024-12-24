package errgroup

import (
	"fmt"
	"net/http"

	"golang.org/x/sync/errgroup"
)

/* JustErrorsExample illustrates the use of a Group in place of a sync.WaitGroup to simplify goroutine counting and error handling. */
func JustErrorsExample(){
	g := new(errgroup.Group)
	urls := []string{
		"http://www.golang.org",
		"http://www.somestupidname.com",
		"http://www.google.com",
	}

	for _, url := range urls {
		url := url // https://golang.org/doc/faq#closures_and_goroutines - Not Required in Go v1.22+
		g.Go(func() error {
			resp, err := http.Get(url)
			if err != nil {
				return err
			}

			body := make([]byte, 9999)
			n, err := resp.Body.Read(body)
			if err != nil {
				return err
			}
			fmt.Printf("Read %v bytes from url %v: %v \n", n, url, string(body))
			resp.Body.Close()

			return nil
		})
	}

	if err := g.Wait(); err == nil {
		fmt.Println("Successfully fetched all URLs.")
	}else{
		fmt.Println("Error Occurred:", err)
	}
}