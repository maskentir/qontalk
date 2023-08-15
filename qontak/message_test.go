package qontak_test

import (
	"errors"
	"testing"

	qontak "github.com/maskentir/qontalk/qontak"
	"github.com/stretchr/testify/assert"
)

func TestSendInteractiveMessage(t *testing.T) {
	tests := []struct {
		name     string
		builder  qontak.SendInteractiveMessage
		strategy *MockRequestStrategy
		expected error
	}{
		{
			name: "SendInteractiveMessage_Success",
			builder: qontak.SendInteractiveMessage{
				RoomID: "room123",
				Type:   "type123",
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
			strategy: &MockRequestStrategy{
				PostResp: map[string]interface{}{
					"result": "success",
				},
			},
			expected: nil,
		},
		{
			name: "SendInteractiveMessage_Failure",
			builder: qontak.SendInteractiveMessage{
				RoomID: "room123",
				Type:   "type123",
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
			strategy: &MockRequestStrategy{
				PostError: errors.New("send interactive message failed"),
			},
			expected: errors.New("send interactive message failed"),
		},
		{
			name:    "SendInteractiveMessage_MissingRequiredFields",
			builder: qontak.SendInteractiveMessage{
				// Missing required fields
			},
			strategy: &MockRequestStrategy{
				PostError: errors.New("send interactive message failed"),
			},
			expected: errors.New("send interactive message failed"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := &qontak.QontakSDK{
				BaseURL:         "https://service-chat.qontak.com/api/open/v1",
				RequestStrategy: tt.strategy,
			}

			err := sdk.SendInteractiveMessage(tt.builder)
			if tt.expected != nil {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.expected.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
