package playground

import (
	"github.com/coderchirag/go-learning/fundamentals/enums"
	"github.com/coderchirag/go-learning/fundamentals/timers"
	libErrors "github.com/coderchirag/go-learning/stdlib/errors"
	libFmt "github.com/coderchirag/go-learning/stdlib/fmt"
	libIter "github.com/coderchirag/go-learning/stdlib/iter"
)

func RunFmtExamples(){
	libFmt.FmtExamples()
}

func RunEnumsExamples(){
	enums.EnumExample()

}

func RunIterExamples(){
	libIter.IterExamples()
}

func RunErrorsExamples(){
	libErrors.ErrorsExamples()
}

func RunTimersExamples() {
	timers.TimersExamples()
}

func Playground() {
	// RunFmtExamples()
	// RunEnumsExamples()
	// RunIterExamples()
	// RunErrorsExamples()
	RunTimersExamples()
}