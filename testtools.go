package idealpostcodes

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
)

// testTools returns a mock HTTP server for unit testing
func testTools(code int, body string) (*httptest.Server, *Client) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, body)
	}))

	transport := &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse(server.URL)
		},
	}

	httpClient := &http.Client{Transport: transport}
	client := NewClient(
		"http://endpoint",
		"the_api_key",
		httpClient,
	)

	return server, client
}
