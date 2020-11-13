package esv

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeRequest(t *testing.T) {
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

	assert.Equal(t, r.Method, "GET")
	assert.Equal(t, r.URL.String(), "http://example.com/one/two?a=true&b=false&c=42&d=foo")
	assert.Equal(t, r.Header.Get("Authorization"), "Token SECRET")
}
