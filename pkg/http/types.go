package http

import "net/url"

type Proxy struct {
	Name     string
	Bind     string
	Protocol string
	Rules    []Rule
	TLS
}

type TLS struct {
	CA   string
	Cert string
	Key  string
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
