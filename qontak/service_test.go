package qontak_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	qontak "github.com/maskentir/qontalk/qontak"
)

type MockRequestStrategy struct {
	PostResp           map[string]interface{}
	PostError          error
	GetResp            map[string]interface{}
	GetError           error
	PutResp            map[string]interface{}
	PutError           error
	PutMultipartResp   map[string]interface{}
	PutMultipartError  error
	PostMultipartResp  map[string]interface{}
	PostMultipartError error
}

func (m *MockRequestStrategy) SetAccessToken(accessToken string) {
	// No need to implement for mock
}

func (m *MockRequestStrategy) Post(
	url string,
	data map[string]interface{},
) (map[string]interface{}, error) {
	if m.PostError != nil {
		return nil, m.PostError
	}
	return m.PostResp, nil
}

func (m *MockRequestStrategy) Get(
	url string,
) (map[string]interface{}, error) {
	if m.GetError != nil {
		return nil, m.GetError
	}
	return m.GetResp, nil
}

func (m *MockRequestStrategy) Put(
	url string,
	data map[string]interface{},
) (map[string]interface{}, error) {
	if m.PutError != nil {
		return nil, m.PutError
	}
	return m.PutResp, nil
}

func (m *MockRequestStrategy) PutMultipart(
	url string,
	formData map[string]interface{},
) (map[string]interface{}, error) {
	if m.PutMultipartError != nil {
		return nil, m.PutMultipartError
	}
	return m.PutMultipartResp, nil
}

func (m *MockRequestStrategy) PostMultipart(
	url string,
	formData map[string]interface{},
) (map[string]interface{}, error) {
	if m.PostMultipartError != nil {
		return nil, m.PostMultipartError
	}
	return m.PostMultipartResp, nil
}

func NewMockRequestStrategy() *MockRequestStrategy {
	return &MockRequestStrategy{}
}

func TestQontakSDK(t *testing.T) {
	tests := []struct {
		name          string
		strategy      qontak.RequestStrategy
		operationFunc func(*qontak.QontakSDK) error
		expectedErr   error
	}{
		{
			name: "Authenticate_Success",
			strategy: &MockRequestStrategy{
				PostResp: map[string]interface{}{
					"access_token": "mockAccessToken",
				},
			},
			operationFunc: func(sdk *qontak.QontakSDK) error {
				return sdk.Authenticate()
			},
			expectedErr: nil,
		},
		{
			name: "Authenticate_Failure",
			strategy: &MockRequestStrategy{
				PostError: errors.New("authentication failed"),
			},
			operationFunc: func(sdk *qontak.QontakSDK) error {
				return sdk.Authenticate()
			},
			expectedErr: errors.New("authentication failed"),
		},
		{
			name: "SendMessageInteractions_Success",
			strategy: &MockRequestStrategy{
				PutMultipartResp: map[string]interface{}{
					"result": "success",
				},
			},
			operationFunc: func(sdk *qontak.QontakSDK) error {
				builder := qontak.SendMessageInteractions{
					ReceiveMessageFromAgent:    true,
					ReceiveMessageFromCustomer: true,
					StatusMessage:              true,
					URL:                        "https://example.com",
				}
				return sdk.SendMessageInteractions(builder)
			},
			expectedErr: nil,
		},
		{
			name: "SendMessageInteractions_Failure",
			strategy: &MockRequestStrategy{
				PutMultipartError: errors.New("send interactions failed"),
			},
			operationFunc: func(sdk *qontak.QontakSDK) error {
				builder := qontak.SendMessageInteractions{
					ReceiveMessageFromAgent:    true,
					ReceiveMessageFromCustomer: true,
					StatusMessage:              true,
					URL:                        "https://example.com",
				}
				return sdk.SendMessageInteractions(builder)
			},
			expectedErr: errors.New("send interactions failed"),
		},
		{
			name: "SendInteractiveMessage_Success",
			strategy: &MockRequestStrategy{
				PostResp: map[string]interface{}{
					"result": "success",
				},
			},
			operationFunc: func(sdk *qontak.QontakSDK) error {
				interactiveData := qontak.NewInteractiveDataBuilder().
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
					Build()

				messageBuilder := qontak.NewSendInteractiveMessageBuilder().
					WithRoomID("room123").
					WithInteractiveData(interactiveData)

				return sdk.SendInteractiveMessage(messageBuilder.Build())
			},
			expectedErr: nil,
		},

		{
			name: "SendInteractiveMessage_Failure",
			strategy: &MockRequestStrategy{
				PostError: errors.New("send interactive message failed"),
			},
			operationFunc: func(sdk *qontak.QontakSDK) error {
				builder := qontak.SendInteractiveMessage{
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
				}
				return sdk.SendInteractiveMessage(builder)
			},
			expectedErr: errors.New("send interactive message failed"),
		},
		{
			name: "SendWhatsAppMessage_Success",
			strategy: &MockRequestStrategy{
				PostMultipartResp: map[string]interface{}{
					"result": "success",
				},
			},
			operationFunc: func(sdk *qontak.QontakSDK) error {
				messageBuilder := qontak.NewWhatsAppMessageBuilder().
					WithRoomID("room123").
					WithMessage("Hello, this is a message!")

				messageParams := messageBuilder.Build()
				return sdk.SendWhatsAppMessage(messageParams)
			},
			expectedErr: nil,
		},
		{
			name: "SendWhatsAppMessage_Failure",
			strategy: &MockRequestStrategy{
				PostMultipartError: errors.New("send WhatsApp message failed"),
			},
			operationFunc: func(sdk *qontak.QontakSDK) error {
				messageBuilder := qontak.NewWhatsAppMessageBuilder().
					WithRoomID("room123").
					WithMessage("Hello, this is a message!")

				messageParams := messageBuilder.Build()
				return sdk.SendWhatsAppMessage(messageParams)
			},
			expectedErr: errors.New("send WhatsApp message failed"),
		},
		{
			name: "SendDirectWhatsAppBroadcast_Success",
			strategy: &MockRequestStrategy{
				PostResp: map[string]interface{}{
					"result": "success",
				},
			},
			operationFunc: func(sdk *qontak.QontakSDK) error {
				broadcastBuilder := qontak.NewDirectWhatsAppBroadcastBuilder().
					WithToName("John Doe").
					WithToNumber("123456789").
					WithMessageTemplateID("template123").
					WithChannelIntegrationID("integration456").
					WithLanguage("en").
					AddDocumentParam("url", "https://example.com/sample.pdf").
					AddDocumentParam("filename", "sample.pdf").
					AddBodyParam("1", "Lorem Ipsum", "customer_name").
					AddButton(qontak.ButtonMessage{Index: "0", Type: "url", Value: "paymentUniqNumber"}).
					Build()

				return sdk.SendDirectWhatsAppBroadcast(broadcastBuilder)
			},
			expectedErr: nil,
		},
		{
			name: "SendDirectWhatsAppBroadcast_Failure",
			strategy: &MockRequestStrategy{
				PostError: errors.New("send direct WhatsApp broadcast failed"),
			},
			operationFunc: func(sdk *qontak.QontakSDK) error {
				broadcastBuilder := qontak.NewDirectWhatsAppBroadcastBuilder().
					WithToName("John Doe").
					WithToNumber("123456789").
					WithMessageTemplateID("template123").
					WithChannelIntegrationID("integration456").
					WithLanguage("en").
					AddDocumentParam("url", "https://example.com/sample.pdf").
					AddDocumentParam("filename", "sample.pdf").
					AddBodyParam("1", "Lorem Ipsum", "customer_name").
					AddButton(qontak.ButtonMessage{Index: "0", Type: "url", Value: "paymentUniqNumber"}).
					Build()

				return sdk.SendDirectWhatsAppBroadcast(broadcastBuilder)
			},
			expectedErr: errors.New("send direct WhatsApp broadcast failed"),
		},
		{
			name: "GetWhatsAppTemplates_Success",
			strategy: &MockRequestStrategy{
				GetResp: map[string]interface{}{
					"template_id": "template123",
				},
			},
			operationFunc: func(sdk *qontak.QontakSDK) error {
				_, err := sdk.GetWhatsAppTemplates()
				return err
			},
			expectedErr: nil,
		},
		{
			name: "GetWhatsAppTemplates_Failure",
			strategy: &MockRequestStrategy{
				GetError: errors.New("get WhatsApp templates failed"),
			},
			operationFunc: func(sdk *qontak.QontakSDK) error {
				_, err := sdk.GetWhatsAppTemplates()
				return err
			},
			expectedErr: errors.New("get WhatsApp templates failed"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := &qontak.QontakSDK{
				BaseURL:         "https://service-chat.qontak.com/api/open/v1",
				RequestStrategy: tt.strategy,
			}

			err := tt.operationFunc(sdk)
			if tt.expectedErr != nil {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestDefaultRequestStrategy(t *testing.T) {
	tests := []struct {
		name            string
		strategy        *qontak.DefaultRequestStrategy
		accessToken     string
		expectedRequest qontak.RequestStrategy
	}{
		{
			name: "StrategyWithAccessToken",
			strategy: &qontak.DefaultRequestStrategy{
				AccessToken: "mockAccessToken",
			},
			accessToken: "mockAccessToken",
			expectedRequest: &qontak.DefaultRequestStrategy{
				AccessToken: "mockAccessToken",
			},
		},
		{
			name: "StrategyWithoutAccessToken",
			strategy: &qontak.DefaultRequestStrategy{
				AccessToken: "",
			},
			accessToken: "",
			expectedRequest: &qontak.DefaultRequestStrategy{
				AccessToken: "",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := &qontak.DefaultRequestStrategy{}
			request.AccessToken = tt.accessToken
			assert.Equal(t, tt.expectedRequest, request)
		})
	}
}
