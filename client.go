package idealpostcodes

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// Client for https://ideal-postcodes.co.uk API
type Client struct {
	host       string
	apiKey     string
	httpClient *http.Client
}

// NewClient returns new Client instance
func NewClient(host string, apiKey string, httpClient *http.Client) *Client {
	return &Client{
		host:       host,
		apiKey:     apiKey,
		httpClient: httpClient,
	}
}

// GetPostcode returns list of addresses matching a postcode
func (c *Client) GetPostcode(postcode string) ([]*Address, error) {
	// Prepare a request
	payload := url.Values{}
	payload.Add("api_key", c.apiKey)
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("%s/postcodes/%s?%s", c.host, postcode, payload.Encode()),
		nil,
	)
	if err != nil {
		return nil, err
	}

	// Make the request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	// Read the response data
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Log the response
	log.Printf("%s", resp.Body)

	// Unmarshall into our struct
	getPostResponse := new(GetPostcodeResponse)
	log.Print(string(data))
	if err := json.Unmarshal(data, getPostResponse); err != nil {
		return nil, err
	}

	// If the status code was not 200, return error with the message from response
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(getPostResponse.Message)
	}

	return getPostResponse.Result, nil
}
