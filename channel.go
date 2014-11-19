package main

import (
	"sync"
)

func NewParallelChannel(src chan *Request, workers int, fn func(<-chan *Request) <-chan *Response) <-chan *Response {
	// spawn worker goroutines
	outputs := make([]<-chan *Response, 0)
	for i := 0; i < workers; i++ {
		outputs = append(outputs, fn(src))
	}

	// Source: http://blog.golang.org/pipelines
	merge := func(cs ...<-chan *Response) <-chan *Response {
		var wg sync.WaitGroup
		out := make(chan *Response)

		output := func(c <-chan *Response) {
			for r := range c {
				out <- r
			}
			wg.Done()
		}
		wg.Add(len(cs))
		for _, c := range cs {
			go output(c)
		}

		go func() {
			wg.Wait()
			close(out)
		}()
		return out
	}

	// join worker goroutines
	return merge(outputs...)
}
