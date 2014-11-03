package main

type RequestQueue struct {
	store []*Request
}

func NewRequestQueue() *RequestQueue {
	return &RequestQueue{store: []*Request{}}
}

func (rq *RequestQueue) Push(req *Request) {
	rq.store = append(rq.store, req)
}

func (rq *RequestQueue) Pop() *Request {
	if len(rq.store) == 0 {
		return nil
	}
	req := rq.store[0]
	rq.store = rq.store[1:]
	return req
}

func (rq *RequestQueue) Len() int {
	return len(rq.store)
}
