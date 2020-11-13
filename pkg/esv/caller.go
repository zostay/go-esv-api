// Package esv provides an SDK for easily interacting with the ESV API located
// at https://api.esv.org/
//
// In order to use this API, you will need an authentication token from them.
package esv

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Client configures the ESV API. Calling New is the preferred way of getting a
// Client object.
type Client struct {
	// BaseURL is the absolute URL to use for API calls. It is set to
	// DefaultBaseURL by default.
	BaseURL *url.URL

	// Client is the http.Client object ot use for making API calls.
	Client *http.Client

	// Token is your authentication token provided to you by esv.org.
	Token string
}

// New will construct a new Client from your authentication token provided to
// you by esv.org.
func New(token string) Client {
	baseURL, err := url.Parse(DefaultBaseURL)
	if err != nil {
		panic(err)
	}

	return Client{
		BaseURL: baseURL,
		Client:  &http.Client{},
		Token:   token,
	}
}

func (c Client) makeRequest(path string, os []Option) (http.Request, error) {
	req := http.Request{
		Method: "GET",
	}

	query := make(url.Values)
	for _, o := range os {
		o.UpdateQuery(query)
	}
	path = path + "?" + query.Encode()

	pathURL, err := url.Parse(path)
	if err != nil {
		return req, fmt.Errorf("error parsing constructed path and query: %w", err)
	}

	requestURL := c.BaseURL.ResolveReference(pathURL)
	req.URL = requestURL

	req.Header = make(http.Header)
	req.Header.Add("Authorization", "Token "+c.Token)

	return req, nil
}

// CallEndpoint is a generic method for making API calls at the endpoint. This
// method is exposed to provide flexibility, but should not normally be used.
func (c Client) CallEndpoint(path string, o []Option, r interface{}) error {
	req, err := c.makeRequest(path, o)
	if err != nil {
		return fmt.Errorf("error building request: %w", err)
	}

	res, err := c.Client.Do(&req)
	if err != nil {
		return fmt.Errorf("error making request: %w", err)
	}

	resBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("error reading response: %w", err)
	}

	err = json.Unmarshal(resBytes, r)
	if err != nil {
		return fmt.Errorf("error decoding response: %w", err)
	}

	return nil
}
