package main

import (
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

func (c *Client) Crawl(url string, depth uint) (*Response, error) {
	if !ValidURL(url) {
		return nil, errors.New("invalid URL")
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "slyther")
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return NewResponse(c.Base, url, resp.Body), nil
}

func (c *Client) Do(url string) {
	queue := NewRequestQueue()
	queue.Push(&Request{url, c.Depth})

	i := 0
	animation := "-\\|/"
	for queue.Len() > 0 {
		fmt.Fprintf(lw, "\rcrawling %v... %c", url, animation[i%len(animation)])
		i++

		req := queue.Pop()
		response, err := c.Crawl(req.Url, req.Depth)
		if err != nil {
			continue
		}
		c.History.Add(response)

		if req.Depth <= 1 {
			continue
		}

		for _, link := range response.Links {
			if !c.History.Has(link) {
				queue.Push(&Request{link, req.Depth - 1})
			}
		}
	}

	fmt.Fprintf(lw, "\rcrawling %v... finished\n", url)
}
