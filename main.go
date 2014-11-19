package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"runtime"
)

const (
	defaultDepth          = 5
	defaultTimeout        = 30
	defaultSourceCapacity = 1024
	defaultOutputFormat   = "text"
	defaultScheme         = "http"
	defaultQuietMode      = false
)

var (
	concurrency = flag.Int("c", runtime.GOMAXPROCS(0), "number of multiple concurrent requests")
	depth       = flag.Uint("depth", defaultDepth, "maximum crawl depth")
	capacity    = flag.Uint("capacity", defaultSourceCapacity, "maximum capacity of source channel")
	format      = flag.String("format", defaultOutputFormat, "output format for sitemap")
	quiet       = flag.Bool("quiet", defaultQuietMode, "quiet mode (don't show progress or error messages)")

	lw io.Writer = os.Stderr // logging writer
	hw io.Writer = os.Stdout // history writer
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

	runtime.GOMAXPROCS(*concurrency)

	if *quiet {
		lw = ioutil.Discard
	}

	client := NewClient(args[0])
	client.Depth = *depth
	client.Capacity = *capacity
	client.Do(client.Base.String())

	switch *format {
	case "html":
		client.History.WriteAsHTML(hw)
	case "json":
		client.History.WriteAsJSON(hw)
	default:
		client.History.WriteAsText(hw)
	}
}
