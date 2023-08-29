package qontak_test

import (
	"testing"

	qontak "github.com/maskentir/qontalk/qontak"
	"github.com/stretchr/testify/assert"
)

func TestBuilders(t *testing.T) {
	tests := []struct {
		name     string
		builder  interface{}
		expected interface{}
	}{
		{
			name: "SendMessageInteractionsBuilder",
			builder: qontak.NewSendMessageInteractionsBuilder().
				WithReceiveMessageFromAgent(true).
				WithReceiveMessageFromCustomer(true).
				WithStatusMessage(true).
				WithURL("https://example.com").
				Build(),
			expected: qontak.SendMessageInteractions{
				ReceiveMessageFromAgent:    true,
				ReceiveMessageFromCustomer: true,
				StatusMessage:              true,
				URL:                        "https://example.com",
			},
		},
		{
			name: "SendInteractiveMessageBuilder",
			builder: qontak.NewSendInteractiveMessageBuilder().
				WithRoomID("room123").
				WithInteractiveData(qontak.NewInteractiveDataBuilder().
					WithHeader(&qontak.InteractiveHeader{
						Format:   "json",
						Text:     "Header Text",
						Link:     "https://example.com",
						Filename: "file.txt",
					}).
					WithBody("Body Text").
					WithButtons([]qontak.Button{
						{ID: "btn1", Title: "Button 1"},
						{ID: "btn2", Title: "Button 2"},
					}).
					Build()).
				Build(),
			expected: qontak.SendInteractiveMessage{
				RoomID: "room123",
				Type:   "string",
				Interactive: qontak.InteractiveData{
					Header: &qontak.InteractiveHeader{
						Format:   "json",
						Text:     "Header Text",
						Link:     "https://example.com",
						Filename: "file.txt",
					},
					Body: "Body Text",
					Buttons: []qontak.Button{
						{ID: "btn1", Title: "Button 1"},
						{ID: "btn2", Title: "Button 2"},
					},
				},
			},
		},
		{
			name: "InteractiveDataBuilder",
			builder: qontak.NewInteractiveDataBuilder().
				WithHeader(&qontak.InteractiveHeader{
					Format:   "json",
					Text:     "Header Text",
					Link:     "https://example.com",
					Filename: "file.txt",
				}).
				WithBody("Body Text").
				WithButtons([]qontak.Button{
					{ID: "btn1", Title: "Button 1"},
					{ID: "btn2", Title: "Button 2"},
				}).
				WithLists(&qontak.InteractiveLists{
					Sections: []qontak.InteractiveSection{
						qontak.NewInteractiveSectionBuilder().
							WithTitle("Section 1").
							WithRows([]qontak.InteractiveRow{
								qontak.NewInteractiveRowBuilder().
									WithID("row1").
									WithTitle("Row 1").
									WithDescription("Description 1").
									Build(),
								qontak.NewInteractiveRowBuilder().
									WithID("row2").
									WithTitle("Row 2").
									WithDescription("Description 2").
									Build(),
							}).
							Build(),
					},
				}).
				Build(),
			expected: qontak.InteractiveData{
				Header: &qontak.InteractiveHeader{
					Format:   "json",
					Text:     "Header Text",
					Link:     "https://example.com",
					Filename: "file.txt",
				},
				Body: "Body Text",
				Buttons: []qontak.Button{
					{ID: "btn1", Title: "Button 1"},
					{ID: "btn2", Title: "Button 2"},
				},
				Lists: &qontak.InteractiveLists{
					Sections: []qontak.InteractiveSection{
						{
							Title: "Section 1",
							Rows: []qontak.InteractiveRow{
								{
									ID:          "row1",
									Title:       "Row 1",
									Description: "Description 1",
								},
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
		},
		{
			name: "InteractiveSectionBuilder",
			builder: qontak.NewInteractiveSectionBuilder().
				WithTitle("Section Title").
				WithRows([]qontak.InteractiveRow{
					qontak.NewInteractiveRowBuilder().
						WithID("row1").
						WithTitle("Row 1").
						WithDescription("Description 1").
						Build(),
					qontak.NewInteractiveRowBuilder().
						WithID("row2").
						WithTitle("Row 2").
						WithDescription("Description 2").
						Build(),
				}).
				Build(),
			expected: qontak.InteractiveSection{
				Title: "Section Title",
				Rows: []qontak.InteractiveRow{
					{
						ID:          "row1",
						Title:       "Row 1",
						Description: "Description 1",
					},
					{
						ID:          "row2",
						Title:       "Row 2",
						Description: "Description 2",
					},
				},
			},
		},
		{
			name: "InteractiveRowBuilder",
			builder: qontak.NewInteractiveRowBuilder().
				WithID("rowID").
				WithTitle("Row Title").
				WithDescription("Row Description").
				Build(),
			expected: qontak.InteractiveRow{
				ID:          "rowID",
				Title:       "Row Title",
				Description: "Row Description",
			},
		},
		{
			name: "DirectWhatsAppBroadcastBuilder",
			builder: qontak.NewDirectWhatsAppBroadcastBuilder().
				WithToName("John Doe").
				WithToNumber("123456789").
				WithMessageTemplateID("template123").
				WithChannelIntegrationID("integration456").
				WithLanguage("en").
				AddHeaderParam("url", "https://example.com/sample.pdf").
				AddHeaderParam("filename", "sample.pdf").
				AddBodyParam("1", "Lorem Ipsum", "customer_name").
				AddButton(qontak.ButtonMessage{Index: "0", Type: "url", Value: "paymentUniqNumber"}).
				Build(),
			expected: qontak.DirectWhatsAppBroadcast{
				ToName:               "John Doe",
				ToNumber:             "123456789",
				MessageTemplateID:    "template123",
				ChannelIntegrationID: "integration456",
				Language:             map[string]string{"code": "en"},
				HeaderParams: []qontak.KeyValue{
					{Key: "url", Value: "https://example.com/sample.pdf"},
					{Key: "filename", Value: "sample.pdf"},
				},
				BodyParams: []qontak.KeyValueText{
					{Key: "1", ValueText: "Lorem Ipsum", Value: "customer_name"},
				},
				Buttons: []qontak.ButtonMessage{
					{Index: "0", Type: "url", Value: "paymentUniqNumber"},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.builder)
		})
	}
}
