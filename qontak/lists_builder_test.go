package qontak_test

import (
	"testing"

	qontak "github.com/maskentir/qontalk/qontak"
	"github.com/stretchr/testify/assert"
)

func TestInteractiveListsBuilder(t *testing.T) {
	tests := []struct {
		name      string
		builder   *qontak.InteractiveListsBuilder
		buildFunc func(*qontak.InteractiveListsBuilder) *qontak.InteractiveLists
		expected  *qontak.InteractiveLists
	}{
		{
			name: "BuildWithButtonAndSections",
			builder: qontak.NewInteractiveListsBuilder().
				WithButton("Button Text").
				WithSections([]qontak.InteractiveSection{
					{
						Title: "Section 1",
						Rows: []qontak.InteractiveRow{
							{
								ID:          "row1",
								Title:       "Row 1",
								Description: "Description 1",
							},
						},
					},
				}),
			buildFunc: func(b *qontak.InteractiveListsBuilder) *qontak.InteractiveLists {
				return b.Build()
			},
			expected: &qontak.InteractiveLists{
				Button: "Button Text",
				Sections: []qontak.InteractiveSection{
					{
						Title: "Section 1",
						Rows: []qontak.InteractiveRow{
							{
								ID:          "row1",
								Title:       "Row 1",
								Description: "Description 1",
							},
						},
					},
				},
			},
		},
		{
			name:    "BuildWithDefaultValues",
			builder: qontak.NewInteractiveListsBuilder(),
			buildFunc: func(b *qontak.InteractiveListsBuilder) *qontak.InteractiveLists {
				return b.Build()
			},
			expected: &qontak.InteractiveLists{},
		},
		{
			name: "BuildWithButtonOnly",
			builder: qontak.NewInteractiveListsBuilder().
				WithButton("Button Text"),
			buildFunc: func(b *qontak.InteractiveListsBuilder) *qontak.InteractiveLists {
				return b.Build()
			},
			expected: &qontak.InteractiveLists{
				Button: "Button Text",
			},
		},
		{
			name: "BuildWithSectionsOnly",
			builder: qontak.NewInteractiveListsBuilder().
				WithSections([]qontak.InteractiveSection{
					{
						Title: "Section 1",
						Rows: []qontak.InteractiveRow{
							{
								ID:          "row1",
								Title:       "Row 1",
								Description: "Description 1",
							},
						},
					},
				}),
			buildFunc: func(b *qontak.InteractiveListsBuilder) *qontak.InteractiveLists {
				return b.Build()
			},
			expected: &qontak.InteractiveLists{
				Sections: []qontak.InteractiveSection{
					{
						Title: "Section 1",
						Rows: []qontak.InteractiveRow{
							{
								ID:          "row1",
								Title:       "Row 1",
								Description: "Description 1",
							},
						},
					},
				},
			},
		},
		{
			name: "BuildWithMultipleSections",
			builder: qontak.NewInteractiveListsBuilder().
				WithSections([]qontak.InteractiveSection{
					{
						Title: "Section 1",
						Rows: []qontak.InteractiveRow{
							{
								ID:          "row1",
								Title:       "Row 1",
								Description: "Description 1",
							},
						},
					},
					{
						Title: "Section 2",
						Rows: []qontak.InteractiveRow{
							{
								ID:          "row2",
								Title:       "Row 2",
								Description: "Description 2",
							},
						},
					},
				}),
			buildFunc: func(b *qontak.InteractiveListsBuilder) *qontak.InteractiveLists {
				return b.Build()
			},
			expected: &qontak.InteractiveLists{
				Sections: []qontak.InteractiveSection{
					{
						Title: "Section 1",
						Rows: []qontak.InteractiveRow{
							{
								ID:          "row1",
								Title:       "Row 1",
								Description: "Description 1",
							},
						},
					},
					{
						Title: "Section 2",
						Rows: []qontak.InteractiveRow{
							{
								ID:          "row2",
								Title:       "Row 2",
								Description: "Description 2",
							},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.buildFunc(tt.builder)
			assert.Equal(t, tt.expected, result)
		})
	}
}
