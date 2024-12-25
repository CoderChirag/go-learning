package playground

import (
	"github.com/coderchirag/go-learning/fundamentals/concurrency/errgroup"
	"github.com/coderchirag/go-learning/fundamentals/enums"
	"github.com/coderchirag/go-learning/fundamentals/timers"
	"github.com/coderchirag/go-learning/fundamentals/worker_pools"
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

func RunWorkerPoolsExamples() {
	worker_pools.WorkerPoolsExample()
}

func RunErrGroupJustErrorsExamples(){
	errgroup.JustErrorsExample()
}

func RunErrorGroupParallelExamples(){
	errgroup.ParallelExample()
}

func RunErrorGroupPipelineExamples(){
	errgroup.PipelineExample()
}


func Playground() {
	// RunFmtExamples()
	// RunEnumsExamples()
	// RunIterExamples()
	// RunErrorsExamples()
	// RunTimersExamples()
	// RunWorkerPoolsExamples()
	// RunErrGroupJustErrorsExamples()
	// RunErrorGroupParallelExamples()
	RunErrorGroupPipelineExamples()
}