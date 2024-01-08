// Package qontak provides a Go SDK for interacting with the Qontak API.
//
// # Documentation
//
// This package offers a Go SDK for accessing the Qontak API. It provides a
// convenient way to authenticate, send message interactions, and send
// interactive messages using the Qontak SDK.
//
// # Overview
//
// The QontakSDKBuilder type is used to build instances of the QontakSDK,
// which is a singleton for accessing the Qontak API. You can use the SDK
// to authenticate, send message interactions, send interactive messages,
// send WhatsApp messages, send Direct WhatsApp Broadcasts, and get WhatsApp
// Templates.
//
// # Authentication
//
// The QontakSDK can be authenticated using the Authenticate method, which
// retrieves an access token for making authenticated API requests.
//
// # Sending Message Interactions
//
// You can use the SendMessageInteractions method to send message interactions,
// including settings for receiving messages from agents or customers, setting
// the status message, and specifying a URL.
//
// # Sending Interactive Messages
//
// The SendInteractiveMessage method allows you to send interactive messages
// to a specified room ID, using interactive data.
//
// # Sending WhatsApp Messages
//
// Use the SendWhatsAppMessage method to send WhatsApp messages to a specified
// room ID with text or images.
//
// # Sending Direct WhatsApp Broadcasts
//
// SendDirectWhatsAppBroadcast enables you to send a direct WhatsApp broadcast
// with custom parameters, including message templates, language settings, and buttons.
//
// # Getting WhatsApp Templates
//
// The GetWhatsAppTemplates method retrieves WhatsApp message templates.
//
// # Customizing Request Strategy
//
// The QontakSDK uses a RequestStrategy interface for sending requests. The
// DefaultRequestStrategy is the default implementation of this interface, but
// you can also set a custom strategy using the SetRequestStrategy method.
//
// # Examples
//
// The following example demonstrates how to use the SDK to send a message
// interaction, an interactive message, a WhatsApp message, a Direct WhatsApp
// Broadcast, and retrieve WhatsApp Templates:
//
//	// Create QontakSDK instance
//	sdkBuilder := qontak.NewQontakSDKBuilder().Build()
//
//	// Authenticate with credentials
//	err := sdkBuilder.Authenticate()
//	if err != nil {
//	    fmt.Println("Authentication failed:", err)
//	    return
//	}
//
//	// Create message interactions builder
//	interactionsBuilder := qontak.NewSendMessageInteractionsBuilder().
//	    WithReceiveMessageFromAgent(true).
//	    WithStatusMessage(true).
//	    WithURL("https://example.com").
//	    Build()
//
//	// Send message interactions
//	err = sdkBuilder.SendMessageInteractions(interactionsBuilder)
//	if err != nil {
//	    fmt.Println("Failed to send interactions:", err)
//	}
//
//	// Create interactive message builder
//	interactiveBuilder := qontak.NewSendInteractiveMessageBuilder().
//	    WithRoomID("room123").
//	    WithInteractiveData(qontak.InteractiveData{
//	        Body: "Hello, World!",
//	        Buttons: []qontak.Button{
//	            {ID: "btn1", Title: "Button 1"},
//	            {ID: "btn2", Title: "Button 2"},
//	        },
//	    }).
//	    Build()
//
//	// Send interactive message
//	err = sdkBuilder.SendInteractiveMessage(interactiveBuilder)
//	if err != nil {
//	    fmt.Println("Failed to send interactive message:", err)
//	}
//
//	// Create WhatsApp message builder
//	whatsappMessageBuilder := qontak.NewWhatsAppMessageBuilder().
//	    WithRoomID("room123").
//	    WithMessage("Hello, this is a message!").
//	    Build()
//
//	// Send WhatsApp message
//	err = sdkBuilder.SendWhatsAppMessage(whatsappMessageBuilder)
//	if err != nil {
//	    fmt.Println("Failed to send WhatsApp message:", err)
//	}
//
//	// Create Direct WhatsApp Broadcast builder
//	directWhatsAppBroadcastBuilder := qontak.NewDirectWhatsAppBroadcastBuilder().
//	    WithToName("John Doe").
//	    WithToNumber("123456789").
//	    WithMessageTemplateID("template123").
//	    WithChannelIntegrationID("integration456").
//	    WithLanguage("en").
//	    AddDocumentParam("url", "https://example.com/sample.pdf").
//	    AddDocumentParam("filename", "sample.pdf").
//	    AddBodyParam("1", "Lorem Ipsum", "customer_name").
//	    AddButton(qontak.ButtonMessage{Index: "0", Type: "url", Value: "paymentUniqNumber"}).
//	    Build()
//
//	// Send Direct WhatsApp Broadcast
//	err = sdkBuilder.SendDirectWhatsAppBroadcast(directWhatsAppBroadcastBuilder)
//	if err != nil {
//	    fmt.Println("Failed to send Direct WhatsApp Broadcast:", err)
//	}
//
//	// Get WhatsApp Templates
//	templates, err := sdkBuilder.GetWhatsAppTemplates()
//	if err != nil {
//	    fmt.Println("Failed to get WhatsApp Templates:", err)
//	} else {
//	    fmt.Println("WhatsApp Templates:", templates)
//	}
package qontak

import (
	"bytes"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
)

// QontakSDKBuilder is a builder to create QontakSDK.
type QontakSDKBuilder struct {
	username     string
	password     string
	grantType    string
	clientID     string
	clientSecret string
}

// NewQontakSDKBuilder creates a new instance of QontakSDKBuilder.
func NewQontakSDKBuilder() *QontakSDKBuilder {
	return &QontakSDKBuilder{}
}

// WithClientCredentials sets client credentials for the builder.
// Example:
// builder.WithClientCredentials("your-username", "your-password", "password", "your-client-id", "your-client-secret")
func (b *QontakSDKBuilder) WithClientCredentials(
	username, password, grantType, clientID, clientSecret string,
) *QontakSDKBuilder {
	b.username = username
	b.password = password
	b.grantType = grantType
	b.clientID = clientID
	b.clientSecret = clientSecret
	return b
}

// Build builds QontakSDK from the builder.
// Example:
// sdk := builder.Build()
func (b *QontakSDKBuilder) Build() *QontakSDK {
	return &QontakSDK{
		BaseURL:         "https://service-chat.qontak.com/api/open/v1",
		Username:        b.username,
		Password:        b.password,
		GrantType:       b.grantType,
		ClientID:        b.clientID,
		ClientSecret:    b.clientSecret,
		RequestStrategy: &DefaultRequestStrategy{},
	}
}

// QontakSDK is a singleton for accessing Qontak API.
type QontakSDK struct {
	BaseURL         string
	Username        string
	Password        string
	GrantType       string
	ClientID        string
	ClientSecret    string
	RequestStrategy RequestStrategy
}

// Authenticate authenticates the SDK with the provided credentials.
// Example:
// err := sdk.Authenticate()
func (sdk *QontakSDK) Authenticate() error {
	authURL := fmt.Sprintf("%s/oauth/token", sdk.BaseURL)

	data := map[string]interface{}{
		"username":      sdk.Username,
		"password":      sdk.Password,
		"grant_type":    sdk.GrantType,
		"client_id":     sdk.ClientID,
		"client_secret": sdk.ClientSecret,
	}

	resp, err := sdk.RequestStrategy.Post(authURL, data)
	fmt.Println(resp)
	if err != nil {
		return err
	}

	accessToken, ok := resp["access_token"].(string)
	if !ok {
		return fmt.Errorf("authentication failed")
	}

	fmt.Println("AccessToken: Bearer", accessToken)
	sdk.RequestStrategy.SetAccessToken(accessToken)
	return nil
}

// SendMessageInteractions sends message interactions.
// Example:
// builder := NewSendMessageInteractionsBuilder().WithReceiveMessageFromAgent(true).WithStatusMessage(true).WithURL("https://example.com")
// err := sdk.SendMessageInteractions(builder.Build())
func (sdk *QontakSDK) SendMessageInteractions(builder SendMessageInteractions) error {
	interactionURL := fmt.Sprintf("%s/message_interactions", sdk.BaseURL)

	data := map[string]interface{}{
		"receive_message_from_agent":    builder.ReceiveMessageFromAgent,
		"receive_message_from_customer": builder.ReceiveMessageFromCustomer,
		"status_message":                builder.StatusMessage,
		"url":                           builder.URL,
	}

	resp, err := sdk.RequestStrategy.PutMultipart(interactionURL, data)
	fmt.Println(resp)
	return err
}

// SendInteractiveMessage sends an interactive message.
// Example:
// builder := NewSendInteractiveMessageBuilder().WithRoomID("room123").WithInteractiveData(interactiveData)
// err := sdk.SendInteractiveMessage(builder.Build())
func (sdk *QontakSDK) SendInteractiveMessage(builder SendInteractiveMessage) error {
	url := fmt.Sprintf("%s/messages/whatsapp/interactive_message/bot", sdk.BaseURL)

	data := map[string]interface{}{
		"room_id":     builder.RoomID,
		"type":        builder.Type,
		"interactive": builder.Interactive,
	}

	resp, err := sdk.RequestStrategy.Post(url, data)
	fmt.Println(resp)
	return err
}

// SendWhatsAppMessage sends a WhatsApp message.
// Example:
// messageBuilder := NewWhatsAppMessageBuilder().
//
//	WithRoomID("room123").
//	WithMessage("Hello, this is a message!").
//	WithImage("path/to/image.png")
//
// messageParams := messageBuilder.Build()
// err := sdk.SendWhatsAppMessage(messageParams)
func (sdk *QontakSDK) SendWhatsAppMessage(params WhatsAppMessage) error {
	url := fmt.Sprintf("%s/messages/whatsapp", sdk.BaseURL)

	formData := map[string]interface{}{
		"room_id": params.RoomID,
		"type":    "text",
		"text":    params.Message,
	}

	resp, err := sdk.RequestStrategy.PostMultipart(url, formData)
	fmt.Println(resp)
	return err
}

// SendDirectWhatsAppBroadcast sends a direct WhatsApp broadcast.
// Example:
// broadcastBuilder := NewDirectWhatsAppBroadcastBuilder().
//
//	WithToName("John Doe").
//	WithToNumber("123456789").
//	WithMessageTemplateID("template123").
//	WithChannelIntegrationID("integration456").
//	WithLanguage("en").
//	AddDocumentParam("url", "https://example.com/sample.pdf").
//	AddDocumentParam("filename", "sample.pdf").
//	AddBodyParam("1", "Lorem Ipsum", "customer_name").
//	AddButton(ButtonMessage{Index: "0", Type: "url", Value: "paymentUniqNumber"}).
//	Build()
//
// err := sdk.SendDirectWhatsAppBroadcast(broadcastBuilder)
func (sdk *QontakSDK) SendDirectWhatsAppBroadcast(params DirectWhatsAppBroadcast) error {
	url := fmt.Sprintf("%s/broadcasts/whatsapp/direct", sdk.BaseURL)

	// Create a data structure to populate the JSON body
	data := map[string]interface{}{
		"to_name":                params.ToName,
		"to_number":              params.ToNumber,
		"message_template_id":    params.MessageTemplateID,
		"channel_integration_id": params.ChannelIntegrationID,
		"language": map[string]interface{}{
			"code": params.Language["code"],
		},
		"parameters": map[string]interface{}{
			"body": convertKeyValueTextToMap(params.BodyParams),
		},
	}

	// Add "document header" only if it exists.
	if len(params.DocumentParams) > 0 {
		data["parameters"].(map[string]interface{})["header"] = map[string]interface{}{
			"format": "DOCUMENT",
			"params": convertKeyValueToMap(params.DocumentParams),
		}
	}

	// Add "image header" only if it exists.
	if len(params.ImageParams) > 0 {
		data["parameters"].(map[string]interface{})["header"] = map[string]interface{}{
			"format": "IMAGE",
			"params": convertKeyValueToMap(params.ImageParams),
		}
	}

	// Add "buttons" only if they exist.
	if len(params.Buttons) > 0 {
		data["parameters"].(map[string]interface{})["buttons"] = convertButtonsToMap(params.Buttons)
	}

	resp, err := sdk.RequestStrategy.Post(url, data)
	fmt.Println(resp)
	return err
}

// GetWhatsAppTemplates mengambil template WhatsApp.
// Example:
// templates, err := sdk.GetWhatsAppTemplates()
func (sdk *QontakSDK) GetWhatsAppTemplates() (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/templates/whatsapp", sdk.BaseURL)

	resp, err := sdk.RequestStrategy.Get(url)
	fmt.Println(resp)
	return resp, err
}

// RequestStrategy is a strategy interface for sending requests
type RequestStrategy interface {
	SetAccessToken(accessToken string)
	// Get sends a Get request with the default strategy.
	// Example:
	// resp, err := drs.Get(url)
	Get(url string) (map[string]interface{}, error)
	// Post sends a POST request with the default strategy.
	// Example:
	// resp, err := drs.Post(url, data)
	Post(url string, data map[string]interface{}) (map[string]interface{}, error)
	// Put sends a PUT request with the default strategy.
	// Example:
	// resp, err := drs.Put(url, data)
	Put(url string, data map[string]interface{}) (map[string]interface{}, error)
	// PutMultipart sends a PUT request with the default strategy.
	// Example:
	// resp, err := drs.PutMultipart(url, formData)
	PutMultipart(
		url string,
		formData map[string]interface{},
	) (map[string]interface{}, error)
	// PostMultipart sends a PUT request with the default strategy.
	// Example:
	// resp, err := drs.PostMultipart(url, formData)
	PostMultipart(
		url string,
		formData map[string]interface{},
	) (map[string]interface{}, error)
}

// DefaultRequestStrategy is the default implementation of RequestStrategy.
type DefaultRequestStrategy struct {
	AccessToken string
}

// SetAccessToken sets the access token in DefaultRequestStrategy.
func (drs *DefaultRequestStrategy) SetAccessToken(accessToken string) {
	drs.AccessToken = accessToken
}

// Get sends a GET request with the default strategy.
// Example:
// resp, err := drs.Get(url)
func (drs *DefaultRequestStrategy) Get(url string) (map[string]interface{}, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	if drs.AccessToken != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", drs.AccessToken))
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var respBody map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		return nil, err
	}

	return respBody, nil
}

// Post sends a POST request with the default strategy.
// Example:
// resp, err := drs.Post(url, data)
func (drs *DefaultRequestStrategy) Post(
	url string,
	data map[string]interface{},
) (map[string]interface{}, error) {
	payloadBytes, _ := json.Marshal(data)
	payload := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	fmt.Println(drs.AccessToken)
	if drs.AccessToken != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", drs.AccessToken))
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var respBody map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		return nil, err
	}

	return respBody, nil
}

// Put sends a PUT request with the default strategy.
// Example:
// resp, err := drs.Put(url, data)
func (drs *DefaultRequestStrategy) Put(
	url string,
	data map[string]interface{},
) (map[string]interface{}, error) {
	payloadBytes, _ := json.Marshal(data)
	payload := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("PUT", url, payload)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	if drs.AccessToken != "" {
		req.Header.Set("Authorization", "Bearer "+drs.AccessToken)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var respBody map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		return nil, err
	}

	return respBody, nil
}

// PutMultipart sends a PUT request with the default strategy.
// Example:
// resp, err := drs.PutMultipart(url, formData)
func (drs *DefaultRequestStrategy) PutMultipart(
	url string,
	formData map[string]interface{},
) (map[string]interface{}, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	for key, value := range formData {
		_ = writer.WriteField(key, fmt.Sprintf("%v", value))
	}

	if err := writer.Close(); err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	if drs.AccessToken != "" {
		req.Header.Set("Authorization", "Bearer "+drs.AccessToken)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var respBody map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		return nil, err
	}

	return respBody, nil
}

// PostMultipart sends a PUT request with the default strategy.
// Example:
// resp, err := drs.PostMultipart(url, formData)
func (drs *DefaultRequestStrategy) PostMultipart(
	url string,
	formData map[string]interface{},
) (map[string]interface{}, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	for key, value := range formData {
		_ = writer.WriteField(key, fmt.Sprintf("%v", value))
	}

	if err := writer.Close(); err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	if drs.AccessToken != "" {
		req.Header.Set("Authorization", "Bearer "+drs.AccessToken)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var respBody map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		return nil, err
	}

	return respBody, nil
}

// SetRequestStrategy sets the request strategy in QontakSDK.
// Example:
// sdk.SetRequestStrategy(&CustomRequestStrategy{})
func (sdk *QontakSDK) SetRequestStrategy(strategy RequestStrategy) {
	sdk.RequestStrategy = strategy
}
