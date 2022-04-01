package main

import (
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/sguessou/greenlight/internal/jsonlog"
)

type testServer struct {
	*httptest.Server
}

func newTestApplication(t *testing.T) *application {
	return &application{
		config: config{
			port: 3007,
			env:  "testing",
		},
		logger: jsonlog.New(os.Stdout, jsonlog.LevelInfo),
	}
}

func newTestServer(t *testing.T, h http.Handler) *testServer {
	ts := httptest.NewServer(h)

	jar, err := cookiejar.New(nil)
	if err != nil {
		t.Fatal(err)
	}

	ts.Client().Jar = jar

	ts.Client().CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}

	return &testServer{ts}
}

func (ts *testServer) get(t *testing.T, urlPath string) (int, http.Header, []byte) {
	rs, err := ts.Client().Get(ts.URL + urlPath)
	if err != nil {
		t.Fatal(err)
	}

	defer rs.Body.Close()
	body, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}

	return rs.StatusCode, rs.Header, body
}

// func (ts *testServer) post(t *testing.T, urlPath string, body string) (int, http.Header, []byte) {
// 	rs, err := ts.Client().Post(ts.URL+urlPath, "application/json", strings.NewReader(body))
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	defer rs.Body.Close()
// 	resp, err := io.ReadAll(rs.Body)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	return rs.StatusCode, rs.Header, resp
// }
