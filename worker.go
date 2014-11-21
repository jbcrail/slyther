package main

import (
	"errors"
	"net/http"
)

func worker(in <-chan *Request) <-chan *Response {
	client := &http.Client{}

	work := func(req *Request) *Response {
		response := &Response{Request: req}
		if !ValidURL(req.Url) {
			response.Error = errors.New("invalid request URL")
			return response
		}
		r, err := http.NewRequest("GET", req.Url, nil)
		if err != nil {
			response.Error = err
			return response
		}
		r.Header.Set("User-Agent", "slyther")
		resp, err := client.Do(r)
		defer resp.Body.Close()
		if err != nil {
			response.Error = err
			return response
		}
		response.ParseHTML(resp.Body)
		return response
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
