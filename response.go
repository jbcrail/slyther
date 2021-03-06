package main

import (
	"io"
	"regexp"

	"code.google.com/p/go.net/html"
)

type Response struct {
	Request *Request
	Links   []string
	Assets  []string
	Error   error `json:"-"`
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

func (resp *Response) ParseHTML(r io.Reader) {
	base, _ := RetrieveBaseURL(defaultScheme, resp.Request.Url)
	re := regexp.MustCompile("^" + base.Scheme + "://" + base.Host)
	linkSet, assetSet := NewSet(), NewSet()
	tokenizer := html.NewTokenizer(r)
	for {
		tokenType := tokenizer.Next()
		if tokenType == html.ErrorToken {
			resp.Links = linkSet.Strings()
			resp.Assets = assetSet.Strings()
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
}

type ResponseByRequestUrl []*Response

func (r ResponseByRequestUrl) Len() int           { return len(r) }
func (r ResponseByRequestUrl) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r ResponseByRequestUrl) Less(i, j int) bool { return r[i].Request.Url < r[j].Request.Url }
