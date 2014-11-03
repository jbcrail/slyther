package main

import (
	"net/url"
	"regexp"
)

func ValidURL(rawurl string) bool {
	u, err := url.Parse(rawurl)
	if err != nil {
		return false
	} else if u.Scheme == "http" || u.Scheme == "https" || u.Scheme == "" {
		return true
	}
	return false
}

func RetrieveBaseURL(scheme string, rawurl string) (*url.URL, error) {
	u, err := url.Parse(rawurl)
	if err != nil {
		return nil, err
	}
	if u.Scheme == "" {
		u.Scheme = scheme
	}
	if u.Host == "" {
		u.Host = u.Path
	}
	u.Path = ""
	return u, nil
}

func ExpandURL(scheme string, host string, rawurl string) (*url.URL, error) {
	u, err := url.Parse(rawurl)
	if err != nil {
		return nil, err
	}
	if u.Scheme == "" {
		u.Scheme = scheme
	}
	if u.Host == "" {
		re := regexp.MustCompile("^.*://")
		u.Host = re.ReplaceAllString(host, "")
	}
	if u.Path == "" {
		u.Path = "/"
	}
	return u, nil
}
