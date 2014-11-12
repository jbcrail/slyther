package main

import (
	"container/heap"
)

type RequestQueue []*Request

func NewRequestQueue() RequestQueue {
	pq := make(RequestQueue, 0)
	heap.Init(&pq)
	return pq
}

func (rq RequestQueue) Len() int {
  return len(rq)
}

func (rq RequestQueue) Less(i, j int) bool {
	return rq[i].Depth < rq[j].Depth
}

func (rq RequestQueue) Swap(i, j int) {
	rq[i], rq[j] = rq[j], rq[i]
}

func (rq *RequestQueue) Push(x interface{}) {
  req := x.(*Request)
	*rq = append(*rq, req)
}

func (rq *RequestQueue) Pop() interface{} {
	old := *rq
	n := len(old)
	req := old[n-1]
	*rq = old[0 : n-1]
	return req
}
