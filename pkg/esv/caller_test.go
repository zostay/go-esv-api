package esv_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/zostay/go-esv-api/pkg/esv"
)

func TestMakeRequest(t *testing.T) {
	t.Parallel()

	testUrl, _ := url.Parse("http://example.com")
	c := esv.Client{
		BaseURL:   testUrl,
		Token:     "SECRET",
		UserAgent: fmt.Sprintf("go-esv-api/%s", esv.Version()),
	}

	r, err := c.MakeRequest(
		"one/two",
		[]esv.Option{
			esv.OptionBool{Name: "a", Value: true},
			esv.OptionBool{Name: "b", Value: false},
			esv.OptionInt{Name: "c", Value: 42},
			esv.OptionString{Name: "d", Value: "foo"},
		},
	)

	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, "GET", r.Method, "GET")
	assert.Equal(t, "http://example.com/one/two?a=true&b=false&c=42&d=foo", r.URL.String())
	assert.Equal(t, "Token SECRET", r.Header.Get("Authorization"))
	assert.Equal(t, "go-esv-api/"+esv.Version(), r.Header.Get("User-Agent"))
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
	c := esv.Client{BaseURL: testUrl, Client: s.Client(), Token: "SECRET"}

	var robj TestResult
	err := c.CallEndpoint(
		context.Background(),
		"zip/zap",
		[]esv.Option{
			esv.OptionBool{Name: "a", Value: true},
			esv.OptionBool{Name: "b", Value: false},
			esv.OptionInt{Name: "c", Value: 42},
			esv.OptionString{Name: "d", Value: "foo"},
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
