package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type posData struct {
	key   string
	value string
}

var theTests = []struct {
	name               string
	url                string
	method             string
	params             []posData
	expectedStatusCode int
}{
	{"home", "/", "GET", []posData{}, http.StatusOK},
	{"about", "/about", "GET", []posData{}, http.StatusOK},
	{"monsters", "/monsters", "GET", []posData{}, http.StatusOK},
	{"weapons", "/weapons", "GET", []posData{}, http.StatusOK},
	{"quests", "/quests", "GET", []posData{}, http.StatusOK},
	{"contact", "/contact", "GET", []posData{}, http.StatusOK},
	{"make-reservation-quest", "/make-reservation-quest", "GET", []posData{}, http.StatusOK},
	{"make-quest", "/make-reservation", "GET", []posData{}, http.StatusOK},
	{"reservation-summary", "/reservation-summary", "GET", []posData{}, http.StatusOK},
}

func TestHandlers(t *testing.T) {
	routes := getRoutes()

	ts := httptest.NewTLSServer(routes)
	defer ts.Close()
	for _, e := range theTests {
		if e.method == "GET" {
			resp, err := ts.Client().Get(ts.URL + e.url)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}
			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("%s: Error in status code. %d", e.name, resp.StatusCode)
			}
		} else {

		}
	}
}
