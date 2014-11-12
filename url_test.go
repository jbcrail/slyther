package main

import (
	"testing"
)

func TestValidURL(t *testing.T) {
	tests := map[string]bool{
		"google.com":                            true,
		"http://google.com/":                    true,
		"https://www.google.com/intl/en/about/": true,
		"/images/icons/product/chrome-48.png":   true,
		"mailto:admin@google.com":               false,
		"javascript:void()":                     false,
	}
	for in, out := range tests {
		if x := ValidURL(in); x != out {
			t.Errorf("ValidURL(%v) = %v, want %v", in, x, out)
		}
	}
}

func TestRetrieveBaseURL(t *testing.T) {
	scheme := "http"
	tests := map[string]string{
		"http://github.com":       "http://github.com/",
		"https://github.com":      "https://github.com/",
		"https://github.com/":     "https://github.com/",
		"http://github.com/about": "http://github.com/",
	}
	for in, out := range tests {
		if url, _ := RetrieveBaseURL(scheme, in); url.String() != out {
			t.Errorf("RetrieveBaseURL(%v, %v) = %v, want %v", scheme, in, url.String(), out)
		}
	}
}

func TestExpandURL(t *testing.T) {
	base, _ := RetrieveBaseURL("http", "acme.com")
	tests := map[string]string{
		"http://github.com":   "http://github.com/",
		"https://github.com":  "https://github.com/",
		"https://github.com/": "https://github.com/",
		"//github.com/about":  "http://github.com/about",
		"/about.html":         "http://acme.com/about.html",
		"about.html":          "http://acme.com/about.html",
	}
	for in, out := range tests {
		if url, _ := ExpandURL(base.Scheme, base.Host, in); url.String() != out {
			t.Errorf("ExpandURL(%v, %v, %v) = %v, want %v", base.Scheme, base.Host, in, url.String(), out)
		}
	}
}
