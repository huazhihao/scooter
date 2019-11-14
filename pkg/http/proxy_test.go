package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestProxy(t *testing.T) {
	backend := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// dummy rely with the url path
		fmt.Fprint(w, r.URL.Path)
	}))
	proxy := &Proxy{}
	frontend := httptest.NewServer(proxy)
	t.Logf("backend: %s", backend.URL)
	t.Logf("scooter: %s", frontend.URL)
	frontendUrl, _ := url.Parse(frontend.URL)
	proxy.UpdateMappings(map[string]string{frontendUrl.Host: backend.URL})

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
