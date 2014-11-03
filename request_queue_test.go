package main

import (
	"testing"
)

func TestEmptyQueue(t *testing.T) {
	rq := NewRequestQueue()
	if rq.Len() != 0 {
		t.Errorf("NewRequestQueue().Len() = %v, want %v", rq.Len(), 0)
	}
}

func TestPushQueue(t *testing.T) {
	rq := NewRequestQueue()
	rq.Push(&Request{"index.html", 2})
	rq.Push(&Request{"about.html", 2})
	rq.Push(&Request{"jobs.html", 1})
	if rq.Len() != 3 {
		t.Errorf("Len() = %v, want %v", rq.Len(), 3)
	}
}

func TestPopQueue(t *testing.T) {
	rq := NewRequestQueue()
	rq.Push(&Request{"index.html", 2})
	rq.Push(&Request{"about.html", 2})
	rq.Push(&Request{"jobs.html", 1})
	rq.Pop()
	rq.Pop()
	if rq.Len() != 1 {
		t.Errorf("Len() = %v, want %v", rq.Len(), 1)
	}
}
