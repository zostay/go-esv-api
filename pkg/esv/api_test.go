package esv_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/zostay/go-esv-api/pkg/esv"
)

type TestHandler func(w http.ResponseWriter, r *http.Request) interface{}

func buildTestClientServer(h TestHandler) (esv.Client, *httptest.Server, *[]*http.Request) {
	reqs := make([]*http.Request, 0, 1)
	var hh http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		result := h(w, r)
		reqs = append(reqs, r)
		bs, _ := json.Marshal(result)
		_, _ = w.Write(bs)
	}

	s := httptest.NewServer(hh)

	testUrl, _ := url.Parse(s.URL)
	c := esv.Client{BaseURL: testUrl, Client: s.Client(), Token: "SECRET"}

	return c, s, &reqs
}

func TextHtmlHandler() TestHandler {
	return func(w http.ResponseWriter, r *http.Request) interface{} {
		return esv.PassageTextResult{
			Query:     "a",
			Canonical: "A",
			Parsed:    [][]esv.Location{{1, 2}, {3, 4}},
			PassageMeta: struct {
				Canonical    string
				ChapterStart []esv.Location
				ChapterEnd   []esv.Location
				PrevVerse    esv.Location
				NextVerse    esv.Location
				PrevChapter  []esv.Location
				NextChapter  []esv.Location
			}{
				Canonical:    "B",
				ChapterStart: []esv.Location{5, 6},
				ChapterEnd:   []esv.Location{7, 8},
				PrevVerse:    9,
				NextVerse:    10,
				PrevChapter:  []esv.Location{11, 12},
				NextChapter:  []esv.Location{13, 14},
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

	p, err := c.PassageText("foo", esv.OptionBool{Name: "bar", Value: true})

	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, 1, len(*reqs))
	assert.Equal(t, "/passage/text?bar=true&q=foo", (*reqs)[0].URL.String())

	assert.Equal(t, "a", p.Query)
}

func TestPassageTextContext(t *testing.T) {
	t.Parallel()

	h := TextHtmlHandler()

	c, s, reqs := buildTestClientServer(h)
	defer s.Close()

	p, err := c.PassageTextContext(context.Background(), "foo", esv.OptionBool{Name: "bar", Value: true})

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

	p, err := c.PassageHtml("foo", esv.OptionBool{Name: "bar", Value: true})

	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, 1, len(*reqs))
	assert.Equal(t, "/passage/html?bar=true&q=foo", (*reqs)[0].URL.String())

	assert.Equal(t, "a", p.Query)
}

func TestPassageHtmlContext(t *testing.T) {
	t.Parallel()

	h := TextHtmlHandler()

	c, s, reqs := buildTestClientServer(h)
	defer s.Close()

	p, err := c.PassageHtmlContext(context.Background(), "foo", esv.OptionBool{Name: "bar", Value: true})

	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, 1, len(*reqs))
	assert.Equal(t, "/passage/html?bar=true&q=foo", (*reqs)[0].URL.String())

	assert.Equal(t, "a", p.Query)
}

func TestWithIncludePassageReferences(t *testing.T) {
	t.Parallel()

	assert.Equal(t, esv.OptionBool{Name: "include-passage-references", Value: true}, esv.WithIncludePassageReferences(true))
}

func TestWithIncludeVerseNumbers(t *testing.T) {
	t.Parallel()

	assert.Equal(t, esv.OptionBool{Name: "include-verse-numbers", Value: true}, esv.WithIncludeVerseNumbers(true))
}

func TestWithIncludeFirstVerseNumbers(t *testing.T) {
	t.Parallel()

	assert.Equal(t, esv.OptionBool{Name: "include-first-verse-numbers", Value: true}, esv.WithIncludeFirstVerseNumbers(true))
}

func TestWithIncludeFootnotes(t *testing.T) {
	t.Parallel()

	assert.Equal(t, esv.OptionBool{Name: "include-footnotes", Value: true}, esv.WithIncludeFootnotes(true))
}

func TestWithIncludeFootnoteBody(t *testing.T) {
	t.Parallel()

	assert.Equal(t, esv.OptionBool{Name: "include-footnote-body", Value: true}, esv.WithIncludeFootnoteBody(true))
}

func TestWithIncludeHeadings(t *testing.T) {
	t.Parallel()

	assert.Equal(t, esv.OptionBool{Name: "include-headings", Value: true}, esv.WithIncludeHeadings(true))
}

func TestWithIncludeShortCopyright(t *testing.T) {
	t.Parallel()

	assert.Equal(t, esv.OptionBool{Name: "include-short-copyright", Value: true}, esv.WithIncludeShortCopyright(true))
}

func TestWithIncludeCopyright(t *testing.T) {
	t.Parallel()

	assert.Equal(t, esv.OptionBool{Name: "include-copyright", Value: true}, esv.WithIncludeCopyright(true))
}

func TestWithIncludePassageHorizontalLines(t *testing.T) {
	t.Parallel()

	assert.Equal(t, esv.OptionBool{Name: "include-passage-horizontal-lines", Value: true}, esv.WithIncludePassageHorizontalLines(true))
}

func TestWithIncludeHeadingHorizontalLines(t *testing.T) {
	t.Parallel()

	assert.Equal(t, esv.OptionBool{Name: "include-heading-horizontal-lines", Value: true}, esv.WithIncludeHeadingHorizontalLines(true))
}

func TestWithHorizontalLineLength(t *testing.T) {
	t.Parallel()

	assert.Equal(t, esv.OptionInt{Name: "horizontal-line-length", Value: 42}, esv.WithHorizontalLineLength(42))
}

func TestWithIncludeSelahs(t *testing.T) {
	t.Parallel()

	assert.Equal(t, esv.OptionBool{Name: "include-selahs", Value: true}, esv.WithIncludeSelahs(true))
}

func TestWithIndentUsing(t *testing.T) {
	t.Parallel()

	assert.Equal(t, esv.OptionString{Name: "indent-using", Value: "foo"}, esv.WithIndentUsing("foo"))
}

func TestWithIndentParagraphs(t *testing.T) {
	t.Parallel()

	assert.Equal(t, esv.OptionInt{Name: "indent-paragraphs", Value: 4}, esv.WithIndentParagraphs(4))
}

func TestWithIndentPoetry(t *testing.T) {
	t.Parallel()

	assert.Equal(t, esv.OptionBool{Name: "indent-poetry", Value: true}, esv.WithIndentPoetry(true))
}

func TestWithIndentPoetryLines(t *testing.T) {
	t.Parallel()

	assert.Equal(t, esv.OptionInt{Name: "indent-poetry-lines", Value: 12}, esv.WithIndentPoetryLines(12))
}

func TestWithIndentDeclares(t *testing.T) {
	t.Parallel()

	assert.Equal(t, esv.OptionInt{Name: "indent-declares", Value: 7}, esv.WithIndentDeclares(7))
}

func TestWithIndentPsalmDoxology(t *testing.T) {
	t.Parallel()

	assert.Equal(t, esv.OptionInt{Name: "indent-psalm-doxology", Value: 3}, esv.WithIndentPsalmDoxology(3))
}

func TestWithLineLength(t *testing.T) {
	t.Parallel()

	assert.Equal(t, esv.OptionInt{Name: "line-length", Value: 80}, esv.WithLineLength(80))
}

func TestWithIncludeCssLink(t *testing.T) {
	t.Parallel()

	assert.Equal(t, esv.OptionBool{Name: "include-css-link", Value: true}, esv.WithIncludeCssLink(true))
}

func TestWithInlineStyles(t *testing.T) {
	t.Parallel()

	assert.Equal(t, esv.OptionBool{Name: "inline-styles", Value: true}, esv.WithInlineStyles(true))
}

func TestWithWrappingDiv(t *testing.T) {
	t.Parallel()

	assert.Equal(t, esv.OptionBool{Name: "wrapping-div", Value: true}, esv.WithWrappingDiv(true))
}

func TestWithDivClasses(t *testing.T) {
	t.Parallel()

	assert.Equal(t, esv.OptionString{Name: "div-classes", Value: "foo"}, esv.WithDivClasses("foo"))
}

func TestWithParagraphTag(t *testing.T) {
	t.Parallel()

	assert.Equal(t, esv.OptionString{Name: "paragraph-tag", Value: "foo"}, esv.WithParagraphTag("foo"))
}

func TestWithIncludeBookTitles(t *testing.T) {
	t.Parallel()

	assert.Equal(t, esv.OptionBool{Name: "include-book-titles", Value: true}, esv.WithIncludeBookTitles(true))
}

func TestWithIncludeVerseAnchors(t *testing.T) {
	t.Parallel()

	assert.Equal(t, esv.OptionBool{Name: "include-verse-anchors", Value: true}, esv.WithIncludeVerseAnchors(true))
}

func TestWithIncludeChapterNumbers(t *testing.T) {
	t.Parallel()

	assert.Equal(t, esv.OptionBool{Name: "include-chapter-numbers", Value: true}, esv.WithIncludeChapterNumbers(true))
}

func TestWithIncludeCrossrefs(t *testing.T) {
	t.Parallel()

	assert.Equal(t, esv.OptionBool{Name: "include-crossrefs", Value: true}, esv.WithIncludeCrossrefs(true))
}

func TestWithIncludeSubheadings(t *testing.T) {
	t.Parallel()

	assert.Equal(t, esv.OptionBool{Name: "include-subheadings", Value: true}, esv.WithIncludeSubheadings(true))
}

func TestWithIncludeSurroundingChapters(t *testing.T) {
	t.Parallel()

	assert.Equal(t, esv.OptionBool{Name: "include-surrounding-chapters", Value: true}, esv.WithIncludeSurroundingChapters(true))
}

func TestWithIncludeSurroundingChaptersBelow(t *testing.T) {
	t.Parallel()

	assert.Equal(t, esv.OptionString{Name: "include-surrounding-chapters-below", Value: "foo"}, esv.WithIncludeSurroundingChaptersBelow("foo"))
}

func TestWithIncludeSurroundingChaptersBelowThreshold(t *testing.T) {
	t.Parallel()

	assert.Equal(t, esv.OptionInt{Name: "include-surrounding-chapters-below-threshold", Value: 42}, esv.WithIncludeSurroundingChaptersBelowThreshold(42))
}

func TestWithLinkUrl(t *testing.T) {
	t.Parallel()

	assert.Equal(t, esv.OptionString{Name: "link-url", Value: "foo"}, esv.WithLinkUrl("foo"))
}

func TestWithCrossrefUrl(t *testing.T) {
	t.Parallel()

	assert.Equal(t, esv.OptionString{Name: "crossref-url", Value: "foo"}, esv.WithCrossrefUrl("foo"))
}

func TestWithPrefaceUrl(t *testing.T) {
	t.Parallel()

	assert.Equal(t, esv.OptionString{Name: "preface-url", Value: "foo"}, esv.WithPrefaceUrl("foo"))
}

func TestWithIncludeAudioLink(t *testing.T) {
	t.Parallel()

	assert.Equal(t, esv.OptionBool{Name: "include-audio-link", Value: true}, esv.WithIncludeAudioLink(true))
}

func TestWithAttachAudioLinkTo(t *testing.T) {
	t.Parallel()

	assert.Equal(t, esv.OptionString{Name: "attach-audio-link-to", Value: "foo"}, esv.WithAttachAudioLinkTo("foo"))
}

func TestWithPageSize(t *testing.T) {
	t.Parallel()

	assert.Equal(t, esv.OptionInt{Name: "page-size", Value: 17}, esv.WithPageSize(17))
}

func TestWithPage(t *testing.T) {
	t.Parallel()

	assert.Equal(t, esv.OptionInt{Name: "page", Value: 2}, esv.WithPage(2))
}
