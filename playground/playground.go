package playground

import (
	"fmt"

	"github.com/coderchirag/go-learning/fundamentals/concurrency/examples/atomic_counters"
	"github.com/coderchirag/go-learning/fundamentals/concurrency/examples/errgroup"
	"github.com/coderchirag/go-learning/fundamentals/concurrency/examples/mutexes"
	"github.com/coderchirag/go-learning/fundamentals/concurrency/examples/rate_limiting"
	"github.com/coderchirag/go-learning/fundamentals/concurrency/examples/stateful_goroutines"
	"github.com/coderchirag/go-learning/fundamentals/concurrency/examples/worker_pools"
	"github.com/coderchirag/go-learning/fundamentals/enums"
	"github.com/coderchirag/go-learning/fundamentals/timers"
	"github.com/coderchirag/go-learning/stdlib/encoding"
	libErrors "github.com/coderchirag/go-learning/stdlib/errors"
	libFmt "github.com/coderchirag/go-learning/stdlib/fmt"
	"github.com/coderchirag/go-learning/stdlib/http"
	libIter "github.com/coderchirag/go-learning/stdlib/iter"
	"github.com/coderchirag/go-learning/stdlib/templates"
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

func RunRateLimitingExamples(){
	rate_limiting.SimpleRateLimiter()
	fmt.Println()
	rate_limiting.BurstyRateLimiter()
}

func RunAtomicCountersExamples(){
	atomic_counters.AtomicCounter()
}

func RunMutexesExamples(){
	mutexes.MutexExample()
}

func RunStatefulGoRoutinesExamples(){
	stateful_goroutines.StatefulGoroutine()
}

func RunTemplatesExamples() {
	templates.Templates()
}

func RunJsonExamples(){
	encoding.JSON()
}

func RunHttpClientExamples(){
	http.Client()
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
	// RunErrorGroupPipelineExamples()
	// RunRateLimitingExamples()
	// RunAtomicCountersExamples()
	// RunMutexesExamples()
	// RunStatefulGoRoutinesExamples()
	// RunTemplatesExamples()
	// RunJsonExamples()
	RunHttpClientExamples()
}