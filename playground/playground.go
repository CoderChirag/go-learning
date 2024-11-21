package playground

import (
	"io"
	"log"
	"os"
	"strings"
)
func Playground() {
	r := strings.NewReader("some random string\n")
	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}
}