package idealpostcodes

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

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

func TestGetPostcode(t *testing.T) {
	data, err := ioutil.ReadFile("_mock_data/get_postcode.json")
	if err != nil {
		log.Fatal(err)
	}

	server, client := testTools(200, string(data))
	defer server.Close()

	addresses, err := client.GetPostcode("ID11QD")

	if assert.Nil(t, err) {
		assert.Equal(t, 7, len(addresses))
		assert.Equal(t, "ID1 1QD", addresses[0].Postcode)
		assert.Equal(t, "LONDON", addresses[0].PostTown)
		assert.Equal(t, "Barons Court Road", addresses[0].Thoroughfare)

		log.Print(addresses[0].UDPRN)

		assert.Equal(t, 25962203, int64(addresses[0].UDPRN))
		assert.Equal(t, "2 Barons Court Road", addresses[0].Line1)
		assert.Equal(t, "", addresses[0].Line2)
		assert.Equal(t, "", addresses[0].Line3)
		assert.Equal(t, "England", addresses[0].Country)
		assert.Equal(t, "Greater London", addresses[0].County)
		assert.Equal(t, "Greater London", addresses[0].TraditionalCounty)
		assert.Equal(t, "Hammersmith and Fulham", addresses[0].District)
		assert.Equal(t, "North End", addresses[0].Ward)
		assert.Equal(t, interface{}(-0.208644362766368), addresses[0].Longitude)
		assert.Equal(t, interface{}(51.4899488390558), addresses[0].Latitude)

		// assert.Equal(t, interface{}(524466), addresses[0].Eastings)
		// assert.Equal(t, interface{}(178299), addresses[0].Northings)
	}
}
