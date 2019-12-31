package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHttpProxyPath(t *testing.T) {
	backend := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// dummy rely with the url path
		fmt.Fprint(w, r.URL.Path)
	}))
	p := &Proxy{}
	frontend := httptest.NewServer(p)
	t.Logf("backend: %s", backend.URL)
	t.Logf("scooter: %s", frontend.URL)
	p.Rules = []Rule{
		Rule{Url: backend.URL},
	}
	p.reload()

	testcases := map[string]string{
		"":         "/",
		"/":        "/",
		"/foo":     "/foo",
		"/foo/":    "/foo/",
		"/foo/bar": "/foo/bar",
	}
	for path, wanted := range testcases {
		resp, err := http.Get(frontend.URL + path)
		if err != nil {
			t.Errorf(err.Error())
		}
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf(err.Error())
		}
		resp.Body.Close()
		s := string(b)
		t.Logf("%s => %s", path, s)
		if s != wanted {
			t.Errorf("resp=%s, want %s", s, wanted)
		}
	}
}

func TestHttpProxyRules(t *testing.T) {
	backend1 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "1")
	}))
	backend2 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "2")
	}))
	p := &Proxy{}
	frontend := httptest.NewServer(p)
	t.Logf("backend1: %s", backend1.URL)
	t.Logf("backend2: %s", backend2.URL)
	t.Logf("scooter: %s", frontend.URL)
	p.Rules = []Rule{
		Rule{Path: "/", Url: backend1.URL},
		Rule{Path: "/v2", Url: backend2.URL},
	}
	p.reload()

	testcases := map[string]string{
		"":        "1",
		"/":       "1",
		"/foo":    "1",
		"/v2":     "2",
		"/v2/foo": "2",
	}
	for path, wanted := range testcases {
		resp, err := http.Get(frontend.URL + path)
		if err != nil {
			t.Errorf(err.Error())
		}
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf(err.Error())
		}
		resp.Body.Close()
		s := string(b)
		t.Logf("%s => %s", path, s)
		if s != wanted {
			t.Errorf("resp=%s, want %s", s, wanted)
		}
	}
}
