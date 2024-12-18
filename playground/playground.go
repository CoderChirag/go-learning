package playground

import (
	"io"
	"log"
	"os"
	"strings"

	libFmt "github.com/coderchirag/go-learning/stdlib/fmt"
)

func Playground() {
	r := strings.NewReader("some random string\n")
	// r := strings.NewReader("some random string\n")

	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}
	// if _, err := io.Copy(os.Stdout, r); err != nil {
	// 	log.Fatal(err)
	// }
	libFmt.FmtExamples()
}