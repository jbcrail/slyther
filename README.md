slyther ![Build Status](https://travis-ci.org/jbcrail/slyther.png)
=======

To install:

    go install github.com/jbcrail/slyther

To test:

    go test github.com/jbcrail/slyther

To run:

    slyther [-depth INT] [-format STRING] url

##### TODO

- output as graphviz
- configure timeout
- guess mimetypes by url; switch to get HEAD of URL?
- save local results w/ timestamps to prevent rude crawling
- obey robots.txt; switch to ignore robots.txt
- handle redirects
- record statistics (memory, time, timeouts, etc)
