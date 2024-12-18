package playground

import (
	"github.com/coderchirag/go-learning/fundamentals/enums"
	libFmt "github.com/coderchirag/go-learning/stdlib/fmt"
)

func Playground() {
	// r := strings.NewReader("some random string\n")

	// if _, err := io.Copy(os.Stdout, r); err != nil {
	// 	log.Fatal(err)
	// }

	libFmt.FmtExamples()
	// libFmt.FmtExamples()
	enums.EnumExample()
}