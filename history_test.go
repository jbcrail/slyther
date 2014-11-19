package main

import (
	"bytes"
	"net/url"
	"reflect"
	"testing"
)

func TestEmptyHistory(t *testing.T) {
	h := NewHistory()
	if h.Len() != 0 {
		t.Errorf("NewHistory().Len() = %v, want %v", h.Len(), 0)
	}
}

func getMockBaseURL() *url.URL {
	base, _ := RetrieveBaseURL("http", "acme.com")
	return base
}

func getMockURLs() []string {
	base := getMockBaseURL()
	url1, _ := ExpandURL(base.Scheme, base.Host, "about.html")
	url2, _ := ExpandURL(base.Scheme, base.Host, "/")
	return []string{url1.String(), url2.String()}
}

func getMockResponses() []*Response {
	urls := getMockURLs()
	reader := bytes.NewReader([]byte{})
	responses := []*Response{}
	responses = append(responses, NewResponse(&Request{urls[0], 0}, reader))
	responses = append(responses, NewResponse(&Request{urls[0], 0}, reader))
	return append(responses, NewResponse(&Request{urls[1], 0}, reader))
}

func TestAddUniqueResponses(t *testing.T) {
	responses := getMockResponses()
	h := NewHistory()
	h.Add(responses[0])
	h.Add(responses[2])
	if h.Len() != 2 {
		t.Errorf("Len() = %v, want %v", h.Len(), 2)
	}
}

func TestAddDuplicateResponses(t *testing.T) {
	responses := getMockResponses()
	h := NewHistory()
	h.Add(responses[0])
	h.Add(responses[1])
	h.Add(responses[2])
	if h.Len() != 2 {
		t.Errorf("Len() = %v, want %v", h.Len(), 2)
	}
}

func TestGetResponse(t *testing.T) {
	responses := getMockResponses()
	h := NewHistory()
	h.Add(responses[0])
	actual := h.Get(responses[0].Request.Url)
	if !reflect.DeepEqual(actual, responses[0]) {
		t.Errorf("Get() = %v, want %v", actual, responses[0])
	}
}

func TestHasResponse(t *testing.T) {
	responses := getMockResponses()
	h := NewHistory()
	h.Add(responses[0])
	actual := h.Has(responses[0].Request.Url)
	expected := true
	if actual != expected {
		t.Errorf("Has() = %v, want %v", actual, expected)
	}
}
