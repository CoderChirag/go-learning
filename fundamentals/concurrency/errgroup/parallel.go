package errgroup

import (
	"context"
	"fmt"
	"os"

	"golang.org/x/sync/errgroup"
)

var (
	Web = fakeSearch("web")
	Image = fakeSearch("image")
	Video = fakeSearch("video")
)

var Google = func(ctx context.Context, query string) ([]Result, error) {
	g, ctx := errgroup.WithContext(ctx)

	searches := []Search{Web, Image, Video}
	results := make([]Result, len(searches))

	for i, search := range searches {
		i, search = i, search // https://golang.org/doc/faq#closures_and_goroutines
		g.Go(func() error {
			result, err := search(ctx, query)
			if err != nil {
				return err
			}
			results[i] = result
			return nil
		})
	}
	
	if err := g.Wait(); err != nil {
		return nil, err
	}

	return results, nil
}

type Result string
type Search func(ctx context.Context, query string) (Result, error)

func fakeSearch(kind string) Search {
	return func(ctx context.Context, query string) (Result, error) {
		return Result(fmt.Sprintf("%s result for %q", kind, query)), nil
	}
}

/* ParallelExample illustrates the use of a Group for synchronizing a simple parallel task */
func ParallelExample(){
	results, err := Google(context.Background(), "golang")

	if err != nil {
		fmt.Println("Error Occurred:", err)
	}

	for i, result := range results {
		fmt.Fprintf(os.Stdout, "Result %d: %s \n", i, string(result))
	}
}