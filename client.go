package main

import (
	"fmt"
	"net/url"
	"runtime"
)

type Client struct {
	Base     *url.URL
	Depth    uint
	Timeout  uint
	Capacity uint
	History  *History
	Errors   uint
}

func NewClient(base string) *Client {
	url, _ := RetrieveBaseURL(defaultScheme, base)
	return &Client{
		Base:     url,
		Depth:    defaultDepth,
		Timeout:  defaultTimeout,
		Capacity: defaultSourceCapacity,
		History:  NewHistory(),
		Errors:   0,
	}
}

func (c *Client) Do(url string) {
	pending := 0

	src := make(chan *Request, c.Capacity)
	sink := NewParallelChannel(src, runtime.GOMAXPROCS(0), worker)

	src <- &Request{Url: url, Depth: 1}
	pending++

	i := 0
	animation := "-\\|/"
	for response := range sink {
		fmt.Fprintf(lw, "\rcrawling %v... %c", url, animation[i%len(animation)])
		i++
		pending--

		if response.Error != nil {
			c.Errors++
			continue
		}

		c.History.Set(response.Request.Url, response)
		for _, link := range response.Links {
			if response.Request.Depth < c.Depth && !c.History.Has(link) {
				c.History.Set(link, nil)
				src <- &Request{Url: link, Depth: response.Request.Depth + 1}
				pending++
			}
		}

		if pending == 0 {
			close(src)
		}
	}

	fmt.Fprintf(lw, "\rcrawling %v... finished\n", url)
}
