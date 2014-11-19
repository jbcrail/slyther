package main

import (
	"net/http"
)

func worker(in <-chan *Request) <-chan *Response {
	client := &http.Client{}

	work := func(req *Request) *Response {
		if !ValidURL(req.Url) {
			return nil
		}
		r, err := http.NewRequest("GET", req.Url, nil)
		if err != nil {
			return nil
		}
		r.Header.Set("User-Agent", "slyther")
		resp, err := client.Do(r)
		if err != nil {
			return nil
		}
		defer resp.Body.Close()
		return NewResponse(req, resp.Body)
	}

	out := make(chan *Response)
	go func() {
		for req := range in {
			out <- work(req)
		}
		close(out)
	}()
	return out
}
