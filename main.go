package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	defaultDepth        = 5
	defaultTimeout      = 30
	defaultOutputFormat = "text"
	defaultScheme       = "http"
)

var (
	depth  = flag.Uint("depth", defaultDepth, "maximum crawl depth")
	format = flag.String("format", defaultOutputFormat, "output format for sitemap")
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: %v [flags] url\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(2)
	}
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		flag.Usage()
	}

	client := NewClient(args[0])
	client.Depth = *depth
	client.Do(client.Base.String())

	switch *format {
	case "html":
		fmt.Println(client.History.ToHTML())
	default:
		fmt.Println(client.History)
	}
}
