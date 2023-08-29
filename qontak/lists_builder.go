package qontak

// InteractiveListsBuilder is a builder for creating interactive message lists.
type InteractiveListsBuilder struct {
	button   string
	sections []InteractiveSection
}

// NewInteractiveListsBuilder creates a new instance of InteractiveListsBuilder.
func NewInteractiveListsBuilder() *InteractiveListsBuilder {
	return &InteractiveListsBuilder{}
}

// WithButton sets the button text for the interactive lists.
func (b *InteractiveListsBuilder) WithButton(button string) *InteractiveListsBuilder {
	b.button = button
	return b
}

// WithSections sets the sections of the interactive lists.
func (b *InteractiveListsBuilder) WithSections(
	sections []InteractiveSection,
) *InteractiveListsBuilder {
	b.sections = sections
	return b
}

// Build constructs an InteractiveLists using the configurations set in the builder.
// Example:
//
//	sectionBuilder := NewInteractiveSectionBuilder().
//	    WithTitle("Section Title").
//	    WithRows([]InteractiveRow{row1, row2})
//	sections := []InteractiveSection{sectionBuilder.Build()}
//
//	listsBuilder := NewInteractiveListsBuilder().
//	    WithButton("View More").
//	    WithSections(sections)
//	lists := listsBuilder.Build()
func (b *InteractiveListsBuilder) Build() *InteractiveLists {
	return &InteractiveLists{
		Button:   b.button,
		Sections: b.sections,
	}
}
