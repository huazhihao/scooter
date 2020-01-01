package http

import "net/url"

type HttpProxy struct {
	Name    string
	Address string
	Rules   []Rule
}

type Rule struct {
	Path    string
	Url     string
	url     *url.URL
	Headers []Header
	Urls    []string
	urls    []*url.URL
}

type Header struct {
	Key   string
	Value string
}
