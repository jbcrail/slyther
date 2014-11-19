package main

import (
	"container/heap"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

type Client struct {
	Base    *url.URL
	Depth   uint
	Timeout uint
	History *History
	client  *http.Client
}

func NewClient(base string) *Client {
	url, _ := RetrieveBaseURL(defaultScheme, base)
	return &Client{Base: url, Depth: defaultDepth, Timeout: defaultTimeout, History: NewHistory(), client: &http.Client{}}
}

func (c *Client) Crawl(req *Request) (*Response, error) {
	if !ValidURL(req.Url) {
		return nil, errors.New("invalid URL")
	}
	r, err := http.NewRequest("GET", req.Url, nil)
	if err != nil {
		return nil, err
	}
	r.Header.Set("User-Agent", "slyther")
	resp, err := c.client.Do(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return NewResponse(req, resp.Body), nil
}

func (c *Client) Do(url string) {
	queue := NewRequestQueue()
	heap.Push(&queue, &Request{Url: url, Depth: 0})

	i := 0
	animation := "-\\|/"
	for queue.Len() > 0 {
		fmt.Fprintf(lw, "\rcrawling %v... %c", url, animation[i%len(animation)])
		i++

		req := heap.Pop(&queue).(*Request)
		if req.Depth >= c.Depth {
			break
		}

		response, err := c.Crawl(req)
		if err != nil {
			continue
		}
		c.History.Add(response)

		for _, link := range response.Links {
			if !c.History.Has(link) {
				heap.Push(&queue, &Request{Url: link, Depth: req.Depth + 1})
			}
		}
	}

	fmt.Fprintf(lw, "\rcrawling %v... finished\n", url)
}
