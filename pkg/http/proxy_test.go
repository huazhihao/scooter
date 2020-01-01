// Copyright © 2019 Hua Zhihao <ihuazhihao@gmail.com>

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHTTPProxyPath(t *testing.T) {
	backend := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// dummy rely with the url path
		fmt.Fprint(w, r.URL.Path)
	}))
	p := &HTTPProxy{}
	frontend := httptest.NewServer(p)
	t.Logf("backend: %s", backend.URL)
	t.Logf("scooter: %s", frontend.URL)
	p.Rules = []Rule{
		Rule{URL: backend.URL},
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

func TestHTTPProxyRules(t *testing.T) {
	backend1 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "1")
	}))
	backend2 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "2")
	}))
	p := &HTTPProxy{}
	frontend := httptest.NewServer(p)
	t.Logf("backend1: %s", backend1.URL)
	t.Logf("backend2: %s", backend2.URL)
	t.Logf("scooter: %s", frontend.URL)
	p.Rules = []Rule{
		Rule{Path: "/", URL: backend1.URL},
		Rule{Path: "/v2", URL: backend2.URL},
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
