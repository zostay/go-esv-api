package esv

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Client struct {
	BaseURL *url.URL
	Client  http.Client
	Token   string
}

func New(token string) Client {
	baseURL, err := url.Parse(DefaultBaseURL)
	if err != nil {
		panic(err)
	}

	return Client{
		BaseURL: baseURL,
		Client:  http.Client{},
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
		return req, err
	}

	requestURL := c.BaseURL.ResolveReference(pathURL)
	req.URL = requestURL

	req.Header = make(http.Header)
	req.Header.Add("Authorization", "Token "+c.Token)

	return req, nil
}

func (c Client) CallEndpoint(path string, o []Option, r interface{}) error {
	req, err := c.makeRequest(path, o)
	if err != nil {
		return err
	}

	res, err := c.Client.Do(&req)
	if err != nil {
		return err
	}

	resBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(resBytes, r)
	if err != nil {
		return err
	}

	return nil
}
