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

func (h *History) Values() []*Response {
	responses := []*Response{}
	for _, response := range h.responses {
		responses = append(responses, response)
	}
	return responses
}

func (h *History) String() string {
	responses := h.Values()
	sort.Sort(ResponseByReferer(responses))

	s := ""
	for _, response := range responses {
		s += response.Referer + "\n"
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
