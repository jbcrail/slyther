package main

import (
	"bytes"
	"sort"
)

type History struct {
	responses map[string]*Response
}

func NewHistory() *History {
	return &History{responses: map[string]*Response{}}
}

func (h *History) Add(resp *Response) {
	h.responses[resp.Referer] = resp
}

func (h *History) Get(referer string) *Response {
	response, _ := h.responses[referer]
	return response
}

func (h *History) Has(referer string) bool {
	_, ok := h.responses[referer]
	return ok
}

func (h *History) Len() int {
	return len(h.responses)
}

func (h *History) Keys() []string {
	keys := []string{}
	for key := range h.responses {
		keys = append(keys, key)
	}
	return keys
}

func (h *History) String() string {
	keys := h.Keys()
	sort.Strings(keys)

	s := ""
	for _, key := range keys {
		response := h.Get(key)
		s += key + "\n"
		for _, link := range response.Links {
			s += "  [link ] " + link + "\n"
		}
		for _, asset := range response.Assets {
			s += "  [asset] " + asset + "\n"
		}
	}
	return string(bytes.TrimSpace([]byte(s)))
}

func (h *History) ToHTML() string {
	return ""
}

func (h *History) ToJSON() string {
	return ""
}
