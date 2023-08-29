package qontak

// InteractiveHeaderBuilder is a builder for creating an interactive message header.
type InteractiveHeaderBuilder struct {
	format   string
	text     string
	link     string
	filename string
}

// NewInteractiveHeaderBuilder creates a new instance of InteractiveHeaderBuilder.
func NewInteractiveHeaderBuilder() *InteractiveHeaderBuilder {
	return &InteractiveHeaderBuilder{}
}

// WithFormat sets the format of the interactive header.
func (b *InteractiveHeaderBuilder) WithFormat(format string) *InteractiveHeaderBuilder {
	b.format = format
	return b
}

// WithText sets the text of the interactive header.
func (b *InteractiveHeaderBuilder) WithText(text string) *InteractiveHeaderBuilder {
	b.text = text
	return b
}

// WithLink sets the link of the interactive header.
func (b *InteractiveHeaderBuilder) WithLink(link string) *InteractiveHeaderBuilder {
	b.link = link
	return b
}

// WithFilename sets the filename of the interactive header.
func (b *InteractiveHeaderBuilder) WithFilename(filename string) *InteractiveHeaderBuilder {
	b.filename = filename
	return b
}

// Build constructs an InteractiveHeader using the configurations set in the builder.
// Example:
//
//	headerBuilder := NewInteractiveHeaderBuilder().
//	    WithFormat("bold").
//	    WithText("Important Header").
//	    WithLink("https://example.com").
//	    WithFilename("header.txt")
//	header := headerBuilder.Build()
func (b *InteractiveHeaderBuilder) Build() *InteractiveHeader {
	return &InteractiveHeader{
		Format:   b.format,
		Text:     b.text,
		Link:     b.link,
		Filename: b.filename,
	}
}
