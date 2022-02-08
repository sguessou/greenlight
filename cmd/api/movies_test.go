package main

import (
	"net/http"
	"testing"
)

func TestPing(t *testing.T) {
	app := newTestApplication(t)

	ts := newTestServer(t, app.routes())
	defer ts.Close()

	code, _, body := ts.get(t, "/ping")
	if code != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, code)
	}

	if string(body) != "OK" {
		t.Errorf("want body equal %q", "OK")
	}
}

// func TestCreateMovieHandlerWithBadJSON(t *testing.T) {
// 	app := newTestApplication(t)

// 	ts := newTestServer(t, app.routes())
// 	defer ts.Close()

// 	t.Run("Happy path", func(t *testing.T) {
// 		jsonBody := `{"title":"Moana","year":2016,"runtime":107,"genres":["animation","adventure"]}`
// 		expected, err := json.MarshalIndent(jsonBody, "", "\t")
// 		if err != nil {
// 			t.Fatal(err)
// 		}
// 		expected = append(expected, '\n')

// 		code, _, body := ts.post(t, "/v1/movies", jsonBody)
// 		if code != http.StatusOK {
// 			t.Errorf("want %d, got %d", http.StatusOK, code)
// 		}

// 		if string(body) != string(expected) {
// 			t.Errorf("want %v, got %v", string(expected), string(body))
// 		}

// 	})
// }
