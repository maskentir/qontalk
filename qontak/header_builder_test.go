package qontak_test

import (
	"testing"

	qontak "github.com/maskentir/qontalk/qontak"
	"github.com/stretchr/testify/assert"
)

func TestInteractiveHeaderBuilder(t *testing.T) {
	tests := []struct {
		name      string
		builder   *qontak.InteractiveHeaderBuilder
		buildFunc func(*qontak.InteractiveHeaderBuilder) *qontak.InteractiveHeader
		expected  *qontak.InteractiveHeader
	}{
		{
			name: "BuildWithAllFields",
			builder: qontak.NewInteractiveHeaderBuilder().
				WithFormat("json").
				WithText("Header Text").
				WithLink("https://example.com").
				WithFilename("file.txt"),
			buildFunc: func(b *qontak.InteractiveHeaderBuilder) *qontak.InteractiveHeader {
				return b.Build()
			},
			expected: &qontak.InteractiveHeader{
				Format:   "json",
				Text:     "Header Text",
				Link:     "https://example.com",
				Filename: "file.txt",
			},
		},
		{
			name:    "BuildWithDefaultValues",
			builder: qontak.NewInteractiveHeaderBuilder(),
			buildFunc: func(b *qontak.InteractiveHeaderBuilder) *qontak.InteractiveHeader {
				return b.Build()
			},
			expected: &qontak.InteractiveHeader{},
		},
		{
			name: "BuildWithFormatAndText",
			builder: qontak.NewInteractiveHeaderBuilder().
				WithFormat("json").
				WithText("Header Text"),
			buildFunc: func(b *qontak.InteractiveHeaderBuilder) *qontak.InteractiveHeader {
				return b.Build()
			},
			expected: &qontak.InteractiveHeader{
				Format: "json",
				Text:   "Header Text",
			},
		},
		{
			name: "BuildWithLinkAndFilename",
			builder: qontak.NewInteractiveHeaderBuilder().
				WithLink("https://example.com").
				WithFilename("file.txt"),
			buildFunc: func(b *qontak.InteractiveHeaderBuilder) *qontak.InteractiveHeader {
				return b.Build()
			},
			expected: &qontak.InteractiveHeader{
				Link:     "https://example.com",
				Filename: "file.txt",
			},
		},
		{
			name: "BuildWithFormatOnly",
			builder: qontak.NewInteractiveHeaderBuilder().
				WithFormat("json"),
			buildFunc: func(b *qontak.InteractiveHeaderBuilder) *qontak.InteractiveHeader {
				return b.Build()
			},
			expected: &qontak.InteractiveHeader{
				Format: "json",
			},
		},
		{
			name: "BuildWithTextOnly",
			builder: qontak.NewInteractiveHeaderBuilder().
				WithText("Header Text"),
			buildFunc: func(b *qontak.InteractiveHeaderBuilder) *qontak.InteractiveHeader {
				return b.Build()
			},
			expected: &qontak.InteractiveHeader{
				Text: "Header Text",
			},
		},
		{
			name: "BuildWithLinkOnly",
			builder: qontak.NewInteractiveHeaderBuilder().
				WithLink("https://example.com"),
			buildFunc: func(b *qontak.InteractiveHeaderBuilder) *qontak.InteractiveHeader {
				return b.Build()
			},
			expected: &qontak.InteractiveHeader{
				Link: "https://example.com",
			},
		},
		{
			name: "BuildWithFilenameOnly",
			builder: qontak.NewInteractiveHeaderBuilder().
				WithFilename("file.txt"),
			buildFunc: func(b *qontak.InteractiveHeaderBuilder) *qontak.InteractiveHeader {
				return b.Build()
			},
			expected: &qontak.InteractiveHeader{
				Filename: "file.txt",
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
