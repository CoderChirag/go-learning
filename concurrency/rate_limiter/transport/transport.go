package transport


type Request[T, K any] struct {
	Args []T
	Fn func(args... T) K
	ResultChan chan K
}

var RequestsQueue = make(chan *Request[any, any], 20)