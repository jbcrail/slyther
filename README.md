Slyther ![Build Status](https://travis-ci.org/jbcrail/slyther.png)
=======

To install:

    go install github.com/jbcrail/slyther

To test:

    go test github.com/jbcrail/slyther

To run:

    slyther [-depth INT] [-format STRING] url

##### TODO

- toggle verbosity (quiet mode?)
- output as graphviz
- configure timeout
- guess mimetypes by url; switch to get HEAD of URL?
- save local results w/ timestamps to prevent rude crawling
- obey robots.txt; switch to ignore robots.txt
- set user agent (ala Mozilla, IE)
