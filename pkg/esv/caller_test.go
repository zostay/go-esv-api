package esv

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeRequest(t *testing.T) {
	t.Parallel()

	testUrl, _ := url.Parse("http://example.com")
	c := Client{BaseURL: testUrl, Token: "SECRET"}

	r, err := c.makeRequest(
		"one/two",
		[]Option{
			OptionBool{"a", true},
			OptionBool{"b", false},
			OptionInt{"c", 42},
			OptionString{"d", "foo"},
		},
	)

	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, "GET", r.Method, "GET")
	assert.Equal(t, "http://example.com/one/two?a=true&b=false&c=42&d=foo", r.URL.String())
	assert.Equal(t, "Token SECRET", r.Header.Get("Authorization"))
}

func TestCallEndpoint(t *testing.T) {
	t.Parallel()

	type TestResult struct {
		Alpha string
		Beta  int
	}

	reqs := make([]*http.Request, 0, 1)
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqs = append(reqs, r)
		bs, _ := json.Marshal(
			TestResult{
				Alpha: "ay",
				Beta:  13,
			},
		)

		_, _ = w.Write(bs)
	}))
	defer s.Close()

	testUrl, _ := url.Parse(s.URL)
	c := Client{BaseURL: testUrl, Client: s.Client(), Token: "SECRET"}

	var robj TestResult
	err := c.CallEndpoint(
		"zip/zap",
		[]Option{
			OptionBool{"a", true},
			OptionBool{"b", false},
			OptionInt{"c", 42},
			OptionString{"d", "foo"},
		},
		&robj,
	)

	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, "ay", robj.Alpha)
	assert.Equal(t, 13, robj.Beta)

	assert.Equal(t, 1, len(reqs))
	assert.Equal(t, "GET", reqs[0].Method)
	assert.Equal(t, "/zip/zap?a=true&b=false&c=42&d=foo", reqs[0].URL.String())
}
