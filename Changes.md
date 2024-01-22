WIP  TBD

 * Breaking Change: `CallEndpoint` now requires a `context.Context` as the first parameter in addition to the other parameters.
 * Breaking Change: `MakeRequest` returns a `*http.Request` instead of `http.Request`.
 * Added the `PassageTextContext` method.
 * Added the `PassageHtmlContext` method.
 * Added the `PassageSearchContext` method.
 * Correction: Fixed errors in documentation that said I had partially implemented `PassageAudio` when I'd actually implemented `PassageSearch`. `PassageAudio` is not yet implemented.

0.1.0  2024-01-22

 * Initial release.
 * Setup a API descriptor file in `tools/gen/esv-api.yaml` that describes the API in a simple format. Then, generate the AIP code using that description using a template in `tools/gen/api.go.tmpl`.
 * Provide support for the `PassageText` API method.
 * The `PassageText` method supports these options:
     * `WithIncludePassageReferences`
     * `WithIncludeVerseNumbers`
     * `WithIncludeFirstVerseNumbers`
     * `WithIncludeFootnotes`
     * `WithIncludeFootnoteBody`
     * `WithIncludeHeadings`
     * `WithIncludeShortCopyright`
     * `WithIncludeCopyright`
     * `WithIncludePassageHorizontalLines`
     * `WithIncludeHeadingHorizontalLines`
     * `WithHorizontalLineLength`
     * `WithIncludeSelahs`
     * `WithIndentUsing`
     * `WithIndentParagraphs`
     * `WithIndentPoetry`
     * `WithIndentPoetryLines`
     * `WithIndentDeclares`
     * `WithIndentPsalmDoxology`
     * `WithLineLength`
 * Provide support for the `PassageHtml`  API method.
 * The `PassageHtml` method supports these options:
     * `WithIncludePassageReferences`
     * `WithIncludeVerseNumbers`
     * `WithIncludeFirstVerseNumbers`
     * `WithIncludeFootnotes`
     * `WithIncludeFootnoteBody`
     * `WithIncludeHeadings`
     * `WithIncludeShortCopyright`
     * `WithIncludeCopyright`
     * `WithIncludeCssLink`
     * `WithInlineStyles`
     * `WithWrappingDiv`
     * `WithDivClasses`
     * `WithParagraphTag`
     * `WithIncludeBookTitles`
     * `WithIncludeVerseAnchors`
     * `WithIncludeChapterNumbers`
     * `WithIncludeCrossrefs`
     * `WithIncludeSubheadings`
     * `WithIncludeSurroundingChapters`
     * `WithIncludeSurroundingChaptersBelow`
     * `WithIncludeChaptersBelowThreshold`
     * `WithLinkUrl`
     * `WithCrossrefUrl`
     * `WithPrefaceUrl`
     * `WithIncludeAUdioLink`
     * `WithAttachAudioLinkTo`
 * Provided support for the `PassageSearch` API method.
 * The `PassageSearch` method supports these options:
     * `WithPageSize`
     * `WithPage`
