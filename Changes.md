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
 * Provide support for the `PassageAudio` API method.
 * The `PassageAudio` method supports these options:
     * `WithPageSize`
     * `WithPage`
