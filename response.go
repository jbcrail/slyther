package main

import (
	"io"
	"net/url"
	"regexp"

	"code.google.com/p/go.net/html"
)

type Response struct {
	Referer string
	Links   []string
	Assets  []string
}

type UrlType int

const (
	LinkUrl = 1 + iota
	AssetUrl
	UnknownUrl
)

func matchUrlTypeAndValue(t html.Token, attribute string, value UrlType) (UrlType, string) {
	for _, attr := range t.Attr {
		if attr.Key == attribute {
			return value, attr.Val
		}
	}
	return UnknownUrl, ""
}

func getUrlTypeAndValue(t html.Token) (UrlType, string) {
	switch t.Data {
	case "a":
		return matchUrlTypeAndValue(t, "href", LinkUrl)
	case "img":
		return matchUrlTypeAndValue(t, "src", AssetUrl)
	case "script":
		return matchUrlTypeAndValue(t, "src", AssetUrl)
	case "link":
		return matchUrlTypeAndValue(t, "href", AssetUrl)
	}
	return UnknownUrl, ""
}

func NewResponse(base *url.URL, referer string, r io.Reader) *Response {
	var response Response
	re := regexp.MustCompile("^" + base.Scheme + "://" + base.Host)
	linkSet, assetSet := NewSet(), NewSet()
	tokenizer := html.NewTokenizer(r)
	for {
		tokenType := tokenizer.Next()
		if tokenType == html.ErrorToken {
			response.Referer = referer
			response.Links = linkSet.Strings()
			response.Assets = assetSet.Strings()
			break
		}
		token := tokenizer.Token()
		if tokenType == html.StartTagToken || tokenType == html.SelfClosingTagToken {
			t, rawurl := getUrlTypeAndValue(token)
			switch t {
			case LinkUrl:
				if ValidURL(rawurl) {
					if url, err := ExpandURL(base.Scheme, base.Host, rawurl); err == nil {
						url.Fragment = ""
						if re.MatchString(url.String()) {
							linkSet.Add(url.String())
						}
					}
				}
			case AssetUrl:
				if ValidURL(rawurl) {
					if url, err := ExpandURL(base.Scheme, base.Host, rawurl); err == nil {
						assetSet.Add(url.String())
					}
				}
			}
		}
	}
	return &response
}

type ResponseByReferer []*Response

func (r ResponseByReferer) Len() int           { return len(r) }
func (r ResponseByReferer) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r ResponseByReferer) Less(i, j int) bool { return r[i].Referer < r[j].Referer }
