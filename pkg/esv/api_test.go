package esv

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestHandler func(w http.ResponseWriter, r *http.Request) interface{}

func buildTestClientServer(h TestHandler) (Client, *httptest.Server, *[]*http.Request) {
	reqs := make([]*http.Request, 0, 1)
	var hh http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		result := h(w, r)
		reqs = append(reqs, r)
		bs, _ := json.Marshal(result)
		_, _ = w.Write(bs)
	}

	s := httptest.NewServer(hh)

	testUrl, _ := url.Parse(s.URL)
	c := Client{BaseURL: testUrl, Client: s.Client(), Token: "SECRET"}

	return c, s, &reqs
}

func TextHtmlHandler() TestHandler {
	return func(w http.ResponseWriter, r *http.Request) interface{} {
		return PassageTextResult{
			Query:     "a",
			Canonical: "A",
			Parsed:    [][]Location{{1, 2}, {3, 4}},
			PassageMeta: struct {
				Canonical    string
				ChapterStart []Location
				ChapterEnd   []Location
				PrevVerse    Location
				NextVerse    Location
				PrevChapter  []Location
				NextChapter  []Location
			}{
				Canonical:    "B",
				ChapterStart: []Location{5, 6},
				ChapterEnd:   []Location{7, 8},
				PrevVerse:    9,
				NextVerse:    10,
				PrevChapter:  []Location{11, 12},
				NextChapter:  []Location{13, 14},
			},
			Passages: []string{"c"},
		}
	}
}

func TestPassageText(t *testing.T) {
	t.Parallel()

	h := TextHtmlHandler()

	c, s, reqs := buildTestClientServer(h)
	defer s.Close()

	p, err := c.PassageText("foo", OptionBool{"bar", true})

	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, 1, len(*reqs))
	assert.Equal(t, "/passage/text?bar=true&q=foo", (*reqs)[0].URL.String())

	assert.Equal(t, "a", p.Query)
}

func TestPassageHtml(t *testing.T) {
	t.Parallel()

	h := TextHtmlHandler()

	c, s, reqs := buildTestClientServer(h)
	defer s.Close()

	p, err := c.PassageHtml("foo", OptionBool{"bar", true})

	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, 1, len(*reqs))
	assert.Equal(t, "/passage/html?bar=true&q=foo", (*reqs)[0].URL.String())

	assert.Equal(t, "a", p.Query)
}
