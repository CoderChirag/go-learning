package errgroup

import (
	"context"
	"crypto/md5"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"golang.org/x/sync/errgroup"
)

/* PipelineExample demonstrates the use of a Group to implement a multi-stage pipeline: a version of the MD5All function with bounded parallelism from https://blog.golang.org/pipelines. */
func PipelineExample(){
	root, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	m, err := md5All(context.Background(), fmt.Sprintf("%s/third-party", root))
	if err != nil {
		log.Fatal(err)
	}

	for path, sum := range m {
		fmt.Printf("%s:\t%x\n", path, sum)
	}
}

type result struct {
	path string
	sum [md5.Size]byte
}

/* MD5All reads all the files in the file tree rooted at root and returns a map from file path to the MD5 sum of the file's contents. If the directory walk fails or any read operation fails, MD5All returns an error. */
func md5All(ctx context.Context, root string) (map[string] [md5.Size]byte, error) {
	g, ctx := errgroup.WithContext(ctx)
	paths := make(chan string)

	g.Go(func() error {
		defer close(paths)
		return filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if d.IsDir() {
				return nil
			}

			select {
				case paths <- path:
				case <-ctx.Done():
					return ctx.Err()	
			}
			return nil
		})
	})

	results := make(chan result)

	g.Go(func() error {
		for path := range paths {
			data, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			select {
				case results <- result{path, md5.Sum(data)}:
				case <-ctx.Done():
					return ctx.Err()
			}
		}
		return nil
	})

	go func() {
		defer close(results)
		g.Wait()
	}()

	m := make(map[string] [md5.Size]byte)
	for result := range results {
		m[result.path] = result.sum
	}
	// Check whether any of the goroutines failed. Since g is accumulating the
	// errors, we don't need to send them (or check for them) in the individual
	// results sent on the channel.
	if err := g.Wait(); err != nil {
		return nil, err
	}

	return m, nil
}

